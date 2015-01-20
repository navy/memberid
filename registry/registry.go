package registry

import (
	"encoding/json"
	"io/ioutil"
)

type Registry struct {
	Members map[string]Member `json:"members"`
	Groups  map[string]Group  `json:"group"`
}

func (r *Registry) ResolveId(id string, from string, to string) string {
	for baseId, u := range r.Members {
		if fromId := u.Id(baseId, from); fromId == id {
			id = u.Id(baseId, to)
			break
		}
	}

	return id
}

func (r *Registry) Ids(groups ...string) []string {
	var ids []string

	if len(groups) > 0 {
		for _, g := range groups {
			if gids, ok := r.Groups[g]; ok {
				ids = append(ids, gids...)
			}
		}
	} else {
		for k, _ := range r.Members {
			ids = append(ids, k)
		}
	}

	return ids
}

type Member map[string]string
type Group []string

func (c Member) Id(id string, typ string) string {
	if v, ok := c[typ]; ok {
		id = v
	}

	return id
}

func LoadConfig(c string) (Registry, error) {
	r := Registry{}

	f, err := ioutil.ReadFile(c)
	if err != nil {
		return r, err
	}

	err = json.Unmarshal(f, &r)
	if err != nil {
		return r, err
	}

	return r, nil
}
