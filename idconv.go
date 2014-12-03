package main

import (
	"encoding/json"
	"io/ioutil"
)

func LoadJson(f string) (map[string]interface{}, error) {
	file, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}

	var config map[string]interface{}
	json.Unmarshal(file, &config)

	return config, nil
}

func ConvertId(id string, config map[string]interface{}, from string, to string) string {
	for baseId, _ := range config {
		data := config[baseId].(map[string]interface{})

		fromId := baseId
		if fromVal, fromFound := data[from]; fromFound {
			fromId = fromVal.(string)
		}

		if fromId == id {
			toId := baseId
			if toVal, toFound := data[to]; toFound {
				toId = toVal.(string)
			}

			id = toId
			break
		}
	}

	return id
}
