package main

import (
	"github.com/Jeffail/gabs"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
)

type Exporter struct {
	status     *Status
	collectors map[string]bool
	metrics    []Metric
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {

	// We're defining our metrics somewhat dynamically and we create the descriptions for our metrics at the point we collect them,
	//  For that reason, we run CollectMetrics() here and in the usual Collect() method
	e.CollectMetrics()

	for _, metric := range e.metrics {
		log.Infoln(metric.Desc)
		ch <- metric.Desc
	}
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {

	e.CollectMetrics()

	for _, metric := range e.metrics {
		if len(metric.Labels) == 0 {
			ch <- prometheus.MustNewConstMetric(
				metric.Desc,
				metric.Type,
				metric.Value,
			)
		} else if len(metric.Labels) == 1 {
			ch <- prometheus.MustNewConstMetric(
				metric.Desc,
				metric.Type,
				metric.Value,
				metric.Labels[0],
			)
		} else if len(metric.Labels) == 2 {
			ch <- prometheus.MustNewConstMetric(
				metric.Desc,
				metric.Type,
				metric.Value,
				metric.Labels[0], metric.Labels[1],
			)
		}
	}
}

func (e *Exporter) CollectMetrics() {
	e.metrics = []Metric{}

	for resource, enabled := range e.collectors {
		if enabled == true {
			e.CollectSummaryMetrics(resource)
		}
	}
}

// CollectResourceDetails collects detailed information about specific resources
func (e *Exporter) CollectDetailedMetrics(resource string, resourceStatus *gabs.Container) error {
	// Remove trailing 's' for prepending to -status-list
	i := resource[:len(resource)-1]

	childMap, err := resourceStatus.Search(i+"-status-list", "status-list-items").ChildrenMap()
	if err != nil {
		return err
	}

	for _, child := range childMap {
		childList, err := child.Children()
		if err != nil {
			return err
		}

		for _, c := range childList {
			var detailStatus *gabs.Container
			var labels []string

			name := c.Path("nameref").Data().(string)

			detailStatus, err = e.status.Get(Path{
				resource: resource,
				name:     name,
			})
			if err != nil {
				return err
			}
			labels = []string{name}

			childMap, err := detailStatus.Search(i+"-status", "status-properties").ChildrenMap()
			if err != nil {
				return err
			}
			err = e.CollectRegister(resource+"_detail", childMap, labels)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// CollectResourceSummary collects and registers metrics for an individual resource summary status URI
func (e *Exporter) CollectSummaryMetrics(resource string) {
	resourceStatus, err := e.status.Get(Path{resource: resource})
	if err != nil {
		return
	}

	// Remove trailing 's' for prepending to -status-list
	i := resource[:len(resource)-1]

	childMap, err := resourceStatus.Search(i+"-status-list", "status-list-summary").ChildrenMap()
	if err != nil {
		return
	}

	err = e.CollectRegister(resource, childMap, []string{})
	if err != nil {
		return
	}

	err = e.CollectDetailedMetrics(resource, resourceStatus)
	if err != nil {
		return
	}

	return
}

// CollectRegister iterates over a JSON structure (in a map[string]*gabs.Container) registering metrics
func (e *Exporter) CollectRegister(resource string, children map[string]*gabs.Container, labels []string) error {
	for key, child := range children {
		if key == "cache-properties" {
			cmap, err := child.ChildrenMap()
			if err != nil {
				return err
			}
			for k, c := range cmap {
				e.AppendMetric(k, c, resource, labels)
			}
		} else if key == "load-properties" {
			cmap, err := child.ChildrenMap()
			if err != nil {
				return err
			}
			for k, c := range cmap {
				if k == "load-detail" {
					cmap, err := c.ChildrenMap()
					if err != nil {
						return err
					}
					for k, c := range cmap {
						e.AppendMetric(k, c, resource, labels)
					}
				} else {
					e.AppendMetric(k, c, resource, labels)
				}
			}
		} else if key == "rate-properties" {
			cmap, err := child.ChildrenMap()
			if err != nil {
				return err
			}
			for k, c := range cmap {
				if k == "rate-detail" {
					cmap, err := c.ChildrenMap()
					if err != nil {
						return err
					}
					for k, c := range cmap {
						e.AppendMetric(k, c, resource, labels)
					}
				} else {
					e.AppendMetric(k, c, resource, labels)
				}
			}
		} else if key != "stands" {
			e.AppendMetric(key, child, resource, labels)
		}
	}
	return nil
}

func (e *Exporter) AppendMetric(name string, child *gabs.Container, subsystem string, labels []string) {
	nm, err := NewMetric(MetricOpts{
		name:      name,
		child:     child,
		subsystem: subsystem,
		labels:    labels,
	})
	if err == nil {
		e.metrics = append(e.metrics, *nm)
	} else {
		log.Errorln(name)
	}
}
