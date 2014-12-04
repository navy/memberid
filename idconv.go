package main

import (
	"encoding/json"
	"io/ioutil"
)

type Registry map[string]User

func (r Registry) ConvertId(id string, from string, to string) string {
	for baseId, u := range r {
		if fromId := u.Id(baseId, from); fromId == id {
			id = u.Id(baseId, to)
			break
		}
	}

	return id
}

type User map[string]string

func (c User) Id(id string, typ string) string {
	if v, ok := c[typ]; ok {
		id = v
	}

	return id
}

func LoadConfig(c string) (Registry, error) {
	f, err := ioutil.ReadFile(c)
	if err != nil {
		return nil, err
	}

	r := Registry{}
	json.Unmarshal(f, &r)

	return r, nil
}
