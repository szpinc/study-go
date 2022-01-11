package util

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// 解析yaml
func ReadYaml(filepath string, out interface{}) error {
	fileBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(fileBytes, out)
}
