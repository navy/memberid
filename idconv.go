package main

import (
	"encoding/json"
	"io/ioutil"
)

func LoadConfig(f string) (map[string]map[string]string, error) {
	file, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}

	config := map[string]map[string]string{}
	json.Unmarshal(file, &config)

	return config, nil
}

func ConvertId(id string, config map[string]map[string]string, from string, to string) string {
	for baseId, _ := range config {
		data := config[baseId]

		fromId := baseId
		if fromVal, fromFound := data[from]; fromFound {
			fromId = fromVal
		}

		if fromId == id {
			toId := baseId
			if toVal, toFound := data[to]; toFound {
				toId = toVal
			}

			id = toId
			break
		}
	}

	return id
}
