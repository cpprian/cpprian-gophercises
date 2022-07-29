package urlhandler

import (
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
)

func InitMux() *http.ServeMux {
	init := http.NewServeMux()
	init.HandleFunc("/", notImplementedPage)
	return init
}

func notImplementedPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not implemented, please be patient"))
}

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
	for i := 0; i < val.Len(); i++ {
		key := val.Index(i).Field(0).String()
		value := val.Index(i).Field(1).String()
		paths[key] = value
	}
	return paths
}
