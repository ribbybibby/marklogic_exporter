package main

import (
	"strconv"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/Jeffail/gabs"
	"strings"
	//"github.com/prometheus/common/log"
)
type Metric struct {
	key 	 string
	value    *gabs.Container
	registry *prometheus.Registry
	subspace string
}

func (metric *Metric) Register() (error) {
	var value float64
	var ok bool

	metric.key = ValidMetricString(metric.key)
	units := ValidMetricString(metric.value.S("units").Data().(string))
	
	if value, ok = metric.value.S("value").Data().(float64); !ok {
		value, _ = strconv.ParseFloat(metric.value.S("value").Data().(string), 64)
	} 

	nm := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: prometheus.BuildFQName(namespace, metric.subspace + "_" + metric.key, units),
		Help: "Placeholder",
	})

	metric.registry.MustRegister(nm)
	nm.Set(value)

	return nil
}

func ValidMetricString(s string) (r string) {
	r = strings.Replace(s, "\"", "", -1)
	r = strings.Replace(r, "-", "_", -1)
	r = strings.Replace(r, "/", "_per_", -1)
	r = strings.Replace(r, "%", "percent", -1)
	return
} 