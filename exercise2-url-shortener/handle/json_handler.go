package urlhandler

import (
	"encoding/json"
	"log"
	"net/http"
)

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
