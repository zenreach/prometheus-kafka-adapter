package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/pkg/errors"
)

type SchemaRegistrySubject struct {
	ID      uint32 `json:"id"`
	Version uint32 `json:"version"`
	Subject string `json:"subject"`
	Schema  string `json:"schema"`
}

func getLatestSubject(registryUrl string, subjectName string) (*SchemaRegistrySubject, error) {
	url, err := url.Parse(registryUrl)
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't parse registry URL.")
	}

	// default to http
	if url.Scheme == "" {
		url.Scheme = "http"
	}

	url.Path = path.Join(url.Path, "subjects", subjectName, "versions", "latest")

	response, err := http.Get(url.String())
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't retrieve schema information from registry.")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't read body from registry response.")
	}

	var subject = new(SchemaRegistrySubject)
	err = json.Unmarshal(body, &subject)

	if err != nil {
		return nil, errors.Wrap(err, "Couldn't unmashall registry response into object.")
	}

	return subject, nil
}
