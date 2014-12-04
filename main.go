package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var Version = "0.1.0"

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
		fmt.Printf("Usage: idconv -c <CONFIG> [OPTION] <ID> \n\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	flag.Parse()

	id = flag.Arg(0)
	if c == "" || id == "" {
		flag.Usage()
	}

	config, err := LoadConfig(c)
	if err != nil {
		log.Fatal(err)
	}

	id = ConvertId(id, config, from, to)

	fmt.Printf("%s\n", id)
}
