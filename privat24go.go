package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "privat24go"
	app.Version = Version
	app.Usage = ""
	app.Author = "Oleg Dolya"
	app.Email = "oleg.dolya@gmail.com"
	app.Commands = Commands

	app.Run(os.Args)
}
