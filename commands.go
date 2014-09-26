package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandLoad,
	commandList,
	commandAdd,
}

var commandLoad = cli.Command{
	Name:  "load",
	Usage: "",
	Description: `
`,
	Action: doLoad,
}

var commandList = cli.Command{
	Name:  "list",
	Usage: "",
	Description: `
`,
	Action: doList,
}

var commandAdd = cli.Command{
	Name:  "add",
	Usage: "",
	Description: `
`,
	Action: doAdd,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doLoad(c *cli.Context) {
}

func doList(c *cli.Context) {
}

func doAdd(c *cli.Context) {
}
