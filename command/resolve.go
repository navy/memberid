package command

import (
	"flag"
	"fmt"

	"github.com/navy/memberid/registry"
)

type ResolveCommand struct {
	From string
	To   string
}

func (c *ResolveCommand) ConfigureFlags(fs *flag.FlagSet) {
	fs.StringVar(&c.From, "from", "", "id-type from")
	fs.StringVar(&c.To, "to", "", "id-type to")
}

func (c *ResolveCommand) Help() string {
	return "[-from=<FROM>] [-to=<TO>] <ID>"
}

func (c *ResolveCommand) Run(fs *flag.FlagSet, r registry.Registry) {
	id := fs.Arg(0)

	id = r.ResolveId(id, c.From, c.To)

	fmt.Printf("%s\n", id)
}
