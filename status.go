package main

import (
	"net/http"
	"github.com/Jeffail/gabs"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"io/ioutil"
	digest "github.com/xinsnake/go-http-digest-auth-client"
)

type Status struct {
	user 	 string
	passwd   string
	uri      string
	registry *prometheus.Registry
}

func (status *Status) Get(p Path) (*gabs.Container, error) {
	var resp *http.Response
	var err error
	var uri string

	uri = p.URI(status.uri)

	dr := digest.NewRequest(status.user, status.passwd, "GET", uri, "")
	if resp, err = dr.Execute(); err != nil {
		return &gabs.Container{}, err
	}

 	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	
	j, err := gabs.ParseJSON(body)

 	return j, nil
}

func (status *Status) Collect(collectors map[string]bool) (error) {
	if collectors["summary"] {
		cluster_status, err := status.Get(Path{})
		if err != nil {
			return err
		}

		for _, a := range [5]string{"forests", "requests", "transactions", "servers", "hosts"} {
			s, _ := cluster_status.Search("local-cluster-status", "status-relations", a + "-status", a + "-status-summary").ChildrenMap()
			status.CollectRegister("summary_" + a, s)
		}
	}

	for v, b := range collectors {
		if b == true && v != "summary" {
			rs, err := status.Get(Path{resource: v})
			if err != nil {
				return err
			}
			i := v[:len(v)-1]
			s, _ := rs.Search(i + "-status-list", "status-list-summary").ChildrenMap()
			status.CollectRegister(v, s)
		}
	}

	return nil
}

func (status *Status) CollectRegister(subspace string, children map[string]*gabs.Container) (error) {
	for key, child := range children {
		if key == "cache-properties" {
			c, err := child.ChildrenMap()
			if err != nil {
				return err
			}
			for k, v := range c {
				metric := Metric{
					key: k,
					value: v,
					registry: status.registry,
					subspace: subspace + "_cache",
				}
				metric.Register()
			}
		} else if key == "load-properties" {
			c, err := child.ChildrenMap()
			if err != nil {
				return err
			}
			for k,v := range c {
				if k == "load-detail" {
					c, err := v.ChildrenMap()
					if err != nil {
						return err
					}
					for l,d := range c {
						metric := Metric{
							key: l,
							value: d,
							registry: status.registry,
							subspace: subspace + "_load_detail",
						}
						metric.Register()					
					}					
				} else {
					metric := Metric{
						key: k,
						value: v,
						registry: status.registry,
						subspace: subspace + "_load",
					}
					metric.Register()
				}
			}
		} else if key == "rate-properties" {
			c, err := child.ChildrenMap()
			if err != nil {
				return err
			}
			for k,v := range c {
				if k == "rate-detail" {
					c, err := v.ChildrenMap()
					if err != nil {
						return err
					}
					for r,d := range c {
						metric := Metric{
							key: r,
							value: d,
							registry: status.registry,
							subspace: subspace + "_load_detail",
						}
						metric.Register()	
					}
				} else {
					metric := Metric{
						key: k,
						value: v,
						registry: status.registry,
						subspace: subspace + "_rate",
					}
					metric.Register()
				}
			}
		} else {
			metric := Metric{
				key: key,
				value: child,
				registry: status.registry,
				subspace: subspace,
			}
			metric.Register()
		}
	}
	return nil
}

