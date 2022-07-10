package urlhandler

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

func InitMux() *http.ServeMux {
	init := http.NewServeMux()
	init.HandleFunc("/", notImplementedPage)
	return init
}

func notImplementedPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not implemented, please be patient"))
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		redirect, ok := pathsToUrls[path]
		if ok {
			http.Redirect(w, r, redirect, http.StatusMovedPermanently)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

type YAMLstruct struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func YAMLtoMap(yml []byte) (map[string]string, error) {

	if f, err := os.OpenFile(string(yml), os.O_RDWR, 0644); err == nil {
		yml, err = io.ReadAll(f)
		if err != nil {
			return nil, err
		}
	}

	var paths []YAMLstruct
	err := yaml.Unmarshal(yml, &paths)
	if err != nil {
		return nil, err
	}
	fmt.Println(paths)

	map_paths := make(map[string]string)
	for _, path := range paths {
		fmt.Println(path)
		map_paths[path.Path] = path.Url
	}
	return map_paths, nil
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	paths, err := YAMLtoMap(yml)
	if err != nil {
		return nil, err
	}

	return MapHandler(paths, fallback), nil
}
