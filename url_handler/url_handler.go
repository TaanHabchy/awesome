package url_handler

import (
	"net/http"
	"strings"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	/*
		Okay, the map handler will recieve a url, then return a http handler, that if present, returns the value to the url's key value in the map
	*/
	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		url := pathsToUrls[req.URL.Path]
		if url != "" {
			http.Redirect(writer, req, url, http.StatusFound)
		} else {
			fallback.ServeHTTP(writer, req)
		}
	})
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
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	key := ""
	value := ""
	pathsToUrls := map[string]string{}
	// parse the string and create a map iterate through the lines with a key and value channel,
	lines := strings.Split(string(yml), "\n")
	// grab the lines
	for _, line := range lines {
		line = strings.TrimSpace(line)
		// if string has prefix -path:. remove that, then set key = to it.
		if strings.HasPrefix(line, "- path: ") {
			key = strings.TrimPrefix(line, "- path: ")
		}
		// then if string has prefix url: , trim and value =.
		if strings.HasPrefix(line, "url: ") {
			value = strings.TrimPrefix(line, "url: ")
		}
		//
		pathsToUrls[key] = value

	}

	return http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
		// check ot see if the requested url is in the map, if not redirect
		if pathsToUrls[req.URL.Path] != "" {
			http.Redirect(writer, req, pathsToUrls[req.URL.Path], http.StatusFound)
		} else {
			fallback.ServeHTTP(writer, req)
		}

	}), nil
}
