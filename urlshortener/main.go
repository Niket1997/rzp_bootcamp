package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fileName := flag.String("jsn", "test.json", "the name of the json file consisting short URLs & their corresponsing long URLs")
	flag.Parse()
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

	yamlHandler, err := YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	jsn, err := ioutil.ReadFile(*fileName)
	if err != nil {
		panic(err)
	}
	jsonHandler, err := JSONHandler([]byte(jsn), yamlHandler)
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
