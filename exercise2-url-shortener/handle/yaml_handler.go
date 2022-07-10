package urlhandler

import (
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

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

func YAMLHandler(data []byte, fallback http.Handler) (http.HandlerFunc, error) {
	yaml_struct := &YamlList{}
	paths, err := yaml_struct.ConvertToMapPaths(data)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return MapHandler(paths, fallback), nil
}
