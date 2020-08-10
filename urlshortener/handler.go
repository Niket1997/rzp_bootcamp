package main

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler to handler map inputs
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		longURL, ok := pathsToUrls[path]
		if ok {
			http.Redirect(w, r, longURL, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// YAMLHandler to handle yaml file
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var urlsPath []urlPath
	err := yaml.Unmarshal(yml, &urlsPath)
	if err != nil {
		return nil, err
	}
	return MapHandler(createMap(urlsPath), fallback), nil
}

// JSONHandler to handle JSON file
func JSONHandler(jsn []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var urlsPath []urlPath
	err := json.Unmarshal(jsn, &urlsPath)
	if err != nil {
		return nil, err
	}

	return MapHandler(createMap(urlsPath), fallback), nil
}

// createMap function to get the map from json or yaml data
func createMap(urlsPath []urlPath) map[string]string {
	pathsToUrls := make(map[string]string, len(urlsPath))
	for _, pathURL := range urlsPath {
		pathsToUrls[pathURL.Path] = pathURL.URL
	}
	return pathsToUrls
}

// urlPath structure
type urlPath struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
