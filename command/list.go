package command

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/navy/memberid/registry"
)

type ListCommand struct {
	Group   string
	To      string
	Shuffle bool
}

func (c *ListCommand) ConfigureFlags(fs *flag.FlagSet) {
	fs.StringVar(&c.Group, "g", "", "group name")
	fs.StringVar(&c.To, "to", "", "id-type to")
	fs.BoolVar(&c.Shuffle, "shuffle", false, "Shuffle ids")
}

func (c *ListCommand) Help() string {
	return "[-g <GROUP>] [-to <TO>]"
}

func (c *ListCommand) Run(fs *flag.FlagSet, r registry.Registry) {
	var ids []string

	if c.Group != "" {
		ids = r.Ids(strings.Split(c.Group, ",")...)
	} else {
		ids = r.Ids()
	}

	if c.Shuffle {
		ids = shuffle(ids)
	}

	if c.To != "" {
		for i, id := range ids {
			ids[i] = r.ResolveId(id, "", c.To)
		}
	}

	fmt.Printf("%v\n", ids)
}

func shuffle(ids []string) []string {
	rand.Seed(time.Now().UTC().UnixNano())
	for i, _ := range ids {
		j := rand.Intn(i + 1)
		ids[i], ids[j] = ids[j], ids[i]
	}

	return ids
}
