package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var (
		id   string
		c    string
		from string
		to   string
	)
	flag.StringVar(&c, "c", "", "config file path (required)")
	flag.StringVar(&from, "from", "", "from type")
	flag.StringVar(&to, "to", "", "to type")

	flag.Usage = func() {
		fmt.Printf("Usage: iconv -c <CONFIG> [OPTION] <ID> \n\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	flag.Parse()

	id = flag.Arg(0)
	if c == "" || id == "" {
		flag.Usage()
	}

	data, err := LoadJson(c)
	if err != nil {
		log.Fatal(err)
	}

	id = ConvertId(id, data, from, to)

	fmt.Printf("%s\n", id)
}

func LoadJson(f string) (map[string]interface{}, error) {
	file, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	json.Unmarshal(file, &data)

	return data, nil
}

func ConvertId(id string, data map[string]interface{}, from string, to string) string {
	for baseId, _udata := range data {
		udata := _udata.(map[string]interface{})

		fromId := baseId
		fromTmp, fromFound := udata[from]
		if fromFound {
			fromId = fromTmp.(string)
		}

		if fromId == id {
			toId := baseId
			toTmp, toFound := udata[to]
			if toFound {
				toId = toTmp.(string)
			}

			id = toId
			break
		}
	}

	return id
}
