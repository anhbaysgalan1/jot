package main

import "github.com/mkideal/cli"

// represents help menu
var help = cli.HelpCommand("display help information")

// represents root command structure
type rootT struct {
    cli.Helper
    Name string `cli:"name" usage:"COMMAND"`
}

// initializes root command context
var root = &cli.Command {
    Desc: "decentralized note-taking at your fingertips",
    Argv: func() interface {} { return new(rootT) },
    Fn: func(ctx *cli.Context) error {
            argv := ctx.Argv().(*rootT)
            ctx.String("Hello, root command, I am %s\n", argv.Name)
            return nil
    },
}

// represents child command structure
type childT struct {
    cli.Helper
    Name string `cli:"name" usage:"COMMAND"`
}

// initializes child command context
var child = &cli.Command {
    Name: "add",
    Desc: "jot down a new note",
    Argv: func() interface{} { return new(childT) },
    Fn: func(ctx *cli.Context) error {
        argv := ctx.Argv().(*childT)
        ctx.String("Hello, child command, I am %s\n", argv.Name)
        return nil
    },
}
