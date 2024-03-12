package urlshort

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...

	return func(w http.ResponseWriter, r *http.Request) {

		// fmt.Printf("the path is: %v", r.URL.Path)

		path := r.URL.Path
		longPath, ok := pathsToUrls[path]

		if !ok {
			fmt.Printf("The path %v is not exist ", path)
			fallback.ServeHTTP(w, r)
			return
		}

		// http.Redirect(w, r, longPath, http.StatusTemporaryRedirect)
		fmt.Fprintf(w, "Redirect to %v", longPath)
		// fmt.Fprintf(w, w.)

	}

}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.

type yamlData struct {
	PATH string
	URL  string
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	var yamlParsedData []yamlData

	err := yaml.Unmarshal([]byte(yml), &yamlParsedData)
	if err != nil {

		return nil, err
	}

	pathsToUrls := map[string]string{}

	for _, yamlEntry := range yamlParsedData {

		pathsToUrls[yamlEntry.PATH] = yamlEntry.URL

	}

	return MapHandler(pathsToUrls, fallback), nil

}
