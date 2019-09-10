package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/sirupsen/logrus"
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
		logrus.WithError(err).Error("Couldn't parse registry URL.")
		return nil, err
	}

	// default to http
	if url.Scheme == "" {
		url.Scheme = "http"
	}

	url.Path = path.Join(url.Path, "subjects", subjectName, "versions", "latest")

	response, err := http.Get(url.String())
	if err != nil {
		logrus.WithError(err).Error("Couldn't retrieve schema information from registry.")
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.WithError(err).Error("Couldn't read body from registry response.")
		return nil, err
	}

	var subject = new(SchemaRegistrySubject)
	err = json.Unmarshal(body, &subject)

	if err != nil {
		logrus.WithError(err).Error("Couldn't unmashall registry response into object.")
		return nil, err
	}

	return subject, nil
}
