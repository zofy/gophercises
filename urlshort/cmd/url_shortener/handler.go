package main

import (
	"encoding/json"
	"net/http"

	"github.com/zofy/urlshort/internal/dbApi"
	yaml "gopkg.in/yaml.v2"
)

type URLMapping struct {
	Path string `yml:"path" json:"path"`
	URL  string `yml:"url" json:"url"`
}

func parseFile(content []byte, format string) ([]URLMapping, error) {
	var urlMappings []URLMapping
	var err error
	if format == "yml" {
		err = yaml.Unmarshal(content, &urlMappings)
	} else if format == "json" {
		err = json.Unmarshal(content, &urlMappings)
	}
	if err != nil {
		return nil, err
	}
	return urlMappings, nil
}

func buildMap(parsedContent []URLMapping) map[string]string {
	pathsURLs := make(map[string]string)
	for _, p := range parsedContent {
		pathsURLs[p.Path] = p.URL
	}
	return pathsURLs
}

// Redirect to address specified in pathsToURLs if found there
// otherwise use fallback http.Hanlder
func MapHandler(pathsToURLs map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if newURL, found := pathsToURLs[req.RequestURI]; found {
			http.Redirect(w, req, newURL, http.StatusSeeOther)
		} else {
			fallback.ServeHTTP(w, req)
		}
	}
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYML, err := parseFile(yml, "yml")
	if err != nil {
		return nil, err
	}
	return MapHandler(buildMap(parsedYML), fallback), nil
}

func JSONHandler(content []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedJSON, err := parseFile(content, "json")
	if err != nil {
		return nil, err
	}
	return MapHandler(buildMap(parsedJSON), fallback), nil
}

func DBHandler(fallback http.Handler) (http.HandlerFunc, error) {
	return func(w http.ResponseWriter, r *http.Request) {
		newURL, err := dbApi.Get(r.RequestURI)
		if err != nil || newURL == "" {
			fallback.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, newURL, http.StatusSeeOther)
		}
	}, nil
}
