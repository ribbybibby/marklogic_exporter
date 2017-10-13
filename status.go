package main

import (
	"net/http"

	"io/ioutil"

	"github.com/Jeffail/gabs"

	digest "github.com/xinsnake/go-http-digest-auth-client"
)

// Status type
type Status struct {
	user   string
	passwd string
	uri    string
}

// Get returns JSON in a *gabs.Container from the path defined by the Path object p
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
