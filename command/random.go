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
	To    string
	Group string
}

func (c *RandomCommand) ConfigureFlags(fs *flag.FlagSet) {
	fs.StringVar(&c.To, "to", "", "id-type to")
	fs.StringVar(&c.Group, "group", "", "group name")
}

func (c *RandomCommand) Help() string {
	return "[-to=<TO>] [-group=<GROUP>]"
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

	rand.Seed(time.Now().Unix())
	id := ids[rand.Intn(len(ids))]

	if c.To != "" {
		id = r.ResolveId(id, "", c.To)
	}

	fmt.Printf("%s\n", id)
}
