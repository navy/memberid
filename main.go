package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/navy/memberid/command"
	"github.com/navy/memberid/registry"
)

var Version = "0.1.1"

type Command interface {
	ConfigureFlags(fs *flag.FlagSet)
	Help() string
	Run(fs *flag.FlagSet, r registry.Registry)
}

func main() {
	var config string

	commands := map[string]Command{
		"resolve": new(command.ResolveCommand),
		"random":  new(command.RandomCommand),
	}

	flag.Usage = func() {
		fmt.Printf("Usage: memberid <command> [-c <CONFIG>] ... \n\n")
		for name, cmd := range commands {
			fmt.Printf("%s: %s\n", name, cmd.Help())
			fs := newFlagSet(name, cmd, &config)
			fs.PrintDefaults()
			fmt.Printf("\n")
		}
		os.Exit(1)
	}

	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()
	}

	name := flag.Arg(0)
	if cmd, ok := commands[name]; ok {

		fs := newFlagSet(name, cmd, &config)

		fs.Parse(flag.Args()[1:])

		r, err := registry.LoadConfig(config)
		if err != nil {
			log.Fatal(err)
		}

		cmd.Run(fs, r)
	} else {
		fmt.Printf("No command found: \"%s\"\n\n", name)
		flag.Usage()
	}
}

func newFlagSet(name string, cmd Command, config *string) *flag.FlagSet {
	fs := flag.NewFlagSet(name, flag.ExitOnError)
	fs.StringVar(config, "c", "memberid.json", "Path to memberid.json file")
	cmd.ConfigureFlags(fs)

	return fs
}
