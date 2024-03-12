package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"urlshort/urlshort"
)

func main() {

	flagFileYamlName := flag.String("yaml", "urls.yaml", "yaml file for urls")
	flag.Parse()

	yamlFile, err := os.Open(*flagFileYamlName)
	if err != nil {
		fmt.Printf("couldn't open the file %v", *flagFileYamlName)
	}

	yaml, err := io.ReadAll(yamlFile)
	if err != nil {
		fmt.Printf("couldn't read the file %v", *flagFileYamlName)
	}
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	// 	yaml := `
	// - path: /urlshort
	//   url: https://github.com/gophercises/urlshort
	// - path: /urlshort-final
	//   url: https://github.com/gophercises/urlshort/tree/solution`
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		fmt.Printf("can't process the request %v", err)
		return
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
