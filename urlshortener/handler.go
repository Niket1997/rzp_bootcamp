package main

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

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

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var urlsPath []urlPath
	err := yaml.Unmarshal(yml, &urlsPath)
	fmt.Println(urlsPath)
	if err != nil {
		return nil, err
	}
	pathsToUrls := make(map[string]string, len(urlsPath))
	for _, pathURL := range urlsPath {
		pathsToUrls[pathURL.Path] = pathURL.URL
	}
	return MapHandler(pathsToUrls, fallback), nil
}

type urlPath struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
