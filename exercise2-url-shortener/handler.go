package urlhandler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"

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

type UrlShortener interface {
	UrlHandler()
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

type Urlconvert interface {
	ConvertToMapPaths(data []byte) (map[string]string, error)
}

type YAMLstruct struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

type YamlList []YAMLstruct

func (yml *YamlList) ConvertToMapPaths(data []byte) (map[string]string, error) {
	ReadDataFromFile(&data)

	err := yaml.Unmarshal(data, yml)
	if err != nil {
		return nil, err
	}

	return ReturnMapPaths(*yml), nil
}

func ReadDataFromFile(data *[]byte) {
	if f, err := os.OpenFile(string(*data), os.O_RDWR, 0644); err == nil {
		*data, err = io.ReadAll(f)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func GetValue(x interface{}) reflect.Value {
	value := reflect.ValueOf(x)

	if value.Kind() == reflect.Ptr {
		return value.Elem()
	}
	return value
}

func ReturnMapPaths(data interface{}) map[string]string {
	val := GetValue(data)
	if val.Kind() != reflect.Slice && val.Kind() != reflect.Array {
		log.Println(val.Kind())
		log.Fatalln("incorrect data type")
		return nil
	}



	paths := make(map[string]string)
	log.Printf("Data -> %q", reflect.ValueOf(val))
	log.Printf("Length of array -> %d", val.Len())

	for i := 0; i < val.Len(); i++ {
		key := val.Index(i).Field(0).String()
		value := val.Index(i).Field(1).String()
		paths[key] = value
	}
	return paths
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
func YAMLHandler(data []byte, fallback http.Handler) (http.HandlerFunc, error) {
	yaml_struct := &YamlList{}
	paths, err := yaml_struct.ConvertToMapPaths(data)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return MapHandler(paths, fallback), nil
}

type JSONstruct struct {
	Path string `json:"path"`
	Url  string `json:"url"`
}

type JsonList []JSONstruct

func (jsn *JsonList) ConvertToMapPaths(data []byte) (map[string]string, error) {
	ReadDataFromFile(&data)

	err := json.Unmarshal(data, jsn)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return ReturnMapPaths(*jsn), nil
}

func JSONHandler(data []byte, fallback http.Handler) (http.HandlerFunc, error) {
	json_struct := &JsonList{}

	paths, err := json_struct.ConvertToMapPaths(data)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return MapHandler(paths, fallback), nil
}
