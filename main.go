package main

import (
	"github.com/ecarter202/inquire"
	"github.com/fatih/color"
)

func main() {
	app := inquire.New().
		Prefix("dadjoke>").
		PrefixColor(color.FgHiRed).
		Commands([]*inquire.Command{
			{
				Label:       "joke",
				Description: "Gets a random joke",
				Handler:     DadJokeHandler,
			},
		},
		)

	app.Run()
}
