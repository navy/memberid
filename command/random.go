package command

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/navy/memberid/registry"
)

type RandomCommand struct {
	Group string
	To    string
}

func (c *RandomCommand) ConfigureFlags(fs *flag.FlagSet) {
	fs.StringVar(&c.Group, "g", "", "group name")
	fs.StringVar(&c.To, "to", "", "id-type to")
}

func (c *RandomCommand) Help() string {
	return "[-g <GROUP>] [-to <TO>]"
}

func (c *RandomCommand) Run(fs *flag.FlagSet, r registry.Registry) {
	var ids []string

	if c.Group != "" {
		ids = r.Ids(strings.Split(c.Group, ",")...)
	} else {
		ids = r.Ids()
	}

	if len(ids) == 0 {
		fmt.Print("No members\n")
		return
	}

	rand.Seed(time.Now().UTC().UnixNano())
	id := ids[rand.Intn(len(ids))]

	if c.To != "" {
		id = r.ResolveId(id, "", c.To)
	}

	fmt.Printf("%s\n", id)
}
