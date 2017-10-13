package main

import (
	"errors"
	"strconv"
	"strings"

	"github.com/Jeffail/gabs"
	"github.com/prometheus/client_golang/prometheus"
)

type Metric struct {
	Type   prometheus.ValueType
	Desc   *prometheus.Desc
	Value  float64
	Labels []string
}

type MetricOpts struct {
	name      string
	child     *gabs.Container
	subsystem string
	labels    []string
}

func NewMetric(opts MetricOpts) (*Metric, error) {
	var desc *prometheus.Desc

	name, err := opts.PrometheusName()
	if err != nil {
		return &Metric{}, err
	}

	if len(opts.labels) == 0 {
		desc = prometheus.NewDesc(
			name,
			"Placeholder",
			nil, nil,
		)
	} else if len(opts.labels) == 1 {
		desc = prometheus.NewDesc(
			name,
			"Placeholder",
			[]string{"name"}, nil,
		)
	} else if len(opts.labels) == 2 {
		desc = prometheus.NewDesc(
			name,
			"Placeholder",
			[]string{"name", "group"}, nil,
		)
	}

	return &Metric{
		Type:   prometheus.GaugeValue,
		Desc:   desc,
		Value:  opts.Value(),
		Labels: opts.labels,
	}, nil
}

func (opts *MetricOpts) PrometheusName() (string, error) {
	units := opts.Units()
	if units == "" {
		return units, errors.New("Empty unit value")
	}

	name := ValidMetricString(opts.name)

	return prometheus.BuildFQName(namespace, opts.subsystem+"_"+name, units), nil
}

// Units comment
func (opts *MetricOpts) Units() (units string) {
	// Only register metrics where we can divine the unit
	if _, ok := opts.child.S("units").Data().(string); ok {
		units = ValidMetricString(opts.child.S("units").Data().(string))
	}
	return
}

// Value converts whatever data type we get from MarkLogic into a float64
func (opts *MetricOpts) Value() float64 {
	var value float64
	var ok bool

	if _, ok = opts.child.S("value").Data().(float64); ok {
		value, _ = opts.child.S("value").Data().(float64)
	} else if _, ok = opts.child.S("value").Data().(string); ok {
		value, _ = strconv.ParseFloat(opts.child.S("value").Data().(string), 64)
	} else if _, ok = opts.child.S("value").Data().(bool); ok {
		if opts.child.S("value").Data().(bool) {
			value = 1
		} else {
			value = 0
		}
	}

	return value
}

// ValidMetricString replaces
func ValidMetricString(s string) (r string) {
	r = strings.Replace(s, "\"", "", -1)
	r = strings.Replace(r, "-", "_", -1)
	r = strings.Replace(r, "/", "_per_", -1)
	r = strings.Replace(r, "%", "percent", -1)
	return
}
