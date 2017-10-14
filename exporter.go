package main

import (
	"strconv"
	"strings"

	"github.com/Jeffail/gabs"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

// Exporter is our custom collector type
type Exporter struct {
	Status       *Status
	Collectors   map[string]bool
	Describing   bool
	DescribeChan *chan<- *prometheus.Desc
	CollectChan  *chan<- prometheus.Metric
}

// Describe all the metrics we're going to export
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {

	// We use the same methods to describe and collect metrics
	//	and we use the Describing bool to tell those methods which we're doing
	e.Describing = true

	// That's also why we capture the channels like this
	e.DescribeChan = &ch

	err := e.Collector()
	if err != nil {
		log.Errorln(err)
	}
}

// Collect the metrics
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	// We use the same methods to describe and collect metrics
	//	and we use the Describing bool to tell those methods which we're doing
	e.Describing = false

	// That's also why we capture the channels like this
	e.CollectChan = &ch

	err := e.Collector()
	if err != nil {
		log.Errorln(err)
	}
}

// Collector is a wrapper function for the methods that do all the actual work
func (e *Exporter) Collector() error {
	var err error

	err = e.CollectSummaryMetrics()
	if err != nil {
		log.Errorln(err)
		return err
	}

	err = e.CollectDetailedMetrics()
	if err != nil {
		log.Errorln(err)
		return err
	}

	return nil
}

// CollectSummaryMetrics returns the basic summary metrics from /v2?view=status
func (e *Exporter) CollectSummaryMetrics() error {

	summaryResources := []string{"forests", "hosts", "requests", "transactions", "servers"}

	summaryStatus, err := e.Status.Get(Path{view: "status"})
	if err != nil {
		return err
	}

	for _, resource := range summaryResources {

		statusProperties, err := summaryStatus.Search("local-cluster-status", "status-relations", resource+"-status", resource+"-status-summary").ChildrenMap()
		if err != nil {
			return err
		}

		e.ParseStatusProperties(statusProperties, resource, map[string]string{})
	}

	return nil
}

// CollectDetailedMetrics iterates through a list of items for each resource retrieved from /v2/$resource.
//	Metrics are retrieved and parsed from /v2/$resource/$name
func (e *Exporter) CollectDetailedMetrics() error {

	for resource, enabled := range e.Collectors {

		if !enabled {
			continue
		}

		resourceShortName := resource[:len(resource)-1]

		resourceStatus, err := e.Status.Get(Path{resource: resource})
		if err != nil {
			log.Errorln(err)
			return err
		}

		resourceList, err := resourceStatus.Search(resourceShortName+"-default-list", "list-items", "list-item").Children()
		if err != nil {
			log.Errorln(err)
			return err
		}

		for _, item := range resourceList {
			var resourceDetails = map[string]string{}

			details, err := item.ChildrenMap()
			if err != nil {
				log.Errorln(err)
				return err
			}

			for key, value := range details {
				if _, ok := value.Data().(string); ok {
					resourceDetails[key] = value.Data().(string)
				}
			}

			path := Path{resource: resource, name: resourceDetails["nameref"], view: "status"}
			labels := map[string]string{
				resourceShortName: resourceDetails["nameref"],
			}

			// You need to supply a group ID for servers
			if resource == "servers" {
				path.group = resourceDetails["groupnameref"]
				labels["group"] = resourceDetails["groupnameref"]
			}

			resourceSummaryStatus, err := e.Status.Get(path)
			if err != nil {
				log.Errorln(err)
				return err
			}

			resourceProperties, err := resourceSummaryStatus.Search(resourceShortName+"-status", "status-properties").ChildrenMap()
			if err != nil {
				log.Errorln(err)
				return err
			}

			e.ParseStatusProperties(resourceProperties, resource, labels)
		}
	}

	return nil
}

// ParseStatusProperties iterates through the status-properties for a given resource and parses values into Prometheus metrics
func (e *Exporter) ParseStatusProperties(statusProperties map[string]*gabs.Container, resource string, labels map[string]string) error {
	for metric, value := range statusProperties {
		if metric == "load-properties" {
			loadProperties, err := value.ChildrenMap()
			if err != nil {
				log.Errorln(err)
				return err
			}
			for metric, value := range loadProperties {
				if metric == "load-detail" {
					loadDetails, err := value.ChildrenMap()
					if err != nil {
						log.Errorln(err)
						return err
					}
					for metric, value := range loadDetails {
						err := e.NewMetric(metric, resource, value, labels)
						if err != nil {
							log.Errorln(err)
							return err
						}
					}
				} else {
					err := e.NewMetric(metric, resource, value, labels)
					if err != nil {
						log.Errorln(err)
						return err
					}
				}
			}
		} else if metric == "rate-properties" {
			rateProperties, err := value.ChildrenMap()
			if err != nil {
				return err
			}
			for metric, value := range rateProperties {
				if metric == "rate-detail" {
					rateDetails, err := value.ChildrenMap()
					if err != nil {
						log.Errorln(err)
						return err
					}
					for metric, value := range rateDetails {
						err := e.NewMetric(metric, resource, value, labels)
						if err != nil {
							log.Errorln(err)
							return err
						}
					}
				} else {
					err := e.NewMetric(metric, resource, value, labels)
					if err != nil {
						log.Errorln(err)
						return err
					}
				}
			}
		} else if metric == "cache-properties" {
			cacheProperties, err := value.ChildrenMap()
			if err != nil {
				log.Errorln(err)
				return err
			}
			for metric, value := range cacheProperties {
				err := e.NewMetric(metric, resource, value, labels)
				if err != nil {
					log.Errorln(err)
					return err
				}
			}
		} else {
			err := e.NewMetric(metric, resource, value, labels)
			if err != nil {
				log.Errorln(err)
				return err
			}
		}
	}

	return nil
}

// NewMetric describes and collects metrics
func (e *Exporter) NewMetric(metric string, subsystem string, value *gabs.Container, labels map[string]string) error {
	var (
		labelNames  []string
		labelValues []string
	)

	for k, v := range labels {
		labelNames = append(labelNames, k)
		labelValues = append(labelValues, v)
	}

	if _, ok := value.S("units").Data().(string); ok {
		name := prometheus.BuildFQName(namespace, subsystem, ValidMetricString(metric)+"_"+ValidMetricString(value.S("units").Data().(string)))
		newDesc := prometheus.NewDesc(
			name,
			"Placeholder",
			labelNames, nil,
		)
		newMetric := prometheus.MustNewConstMetric(
			newDesc,
			prometheus.GaugeValue,
			ValidMetricValue(value),
			labelValues...,
		)
		if e.Describing {
			*e.DescribeChan <- newDesc
		} else {
			*e.CollectChan <- newMetric
		}
	}
	return nil
}

// ValidMetricString replaces non-Prometheus friendly characters
func ValidMetricString(s string) string {
	var r string
	r = strings.Replace(s, "\"", "", -1)
	r = strings.Replace(r, "-", "_", -1)
	r = strings.Replace(r, "/", "_per_", -1)
	r = strings.Replace(r, "%", "percent", -1)
	return r
}

// ValidMetricValue converts whatever data type we get from MarkLogic into a float64
func ValidMetricValue(gab *gabs.Container) float64 {
	var value float64
	var ok bool

	if _, ok = gab.S("value").Data().(float64); ok {
		value, _ = gab.S("value").Data().(float64)
	} else if _, ok = gab.S("value").Data().(string); ok {
		value, _ = strconv.ParseFloat(gab.S("value").Data().(string), 64)
	} else if _, ok = gab.S("value").Data().(bool); ok {
		if gab.S("value").Data().(bool) {
			value = 1
		} else {
			value = 0
		}
	}

	return value
}
