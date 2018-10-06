package main

import (
    "fmt"
    "os"

    "github.com/mkideal/cli"
    "github.com/ipfs/go-ipfs-api"
)

// ipfs global bindings
var sh *shell.Shell
var ncalls int


func main() {

    var cli_root = cli.Root(root,
                            cli.Tree(help),
                            cli.Tree(child))

    var err = cli_root.Run(os.Args[1:]);
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }

    // initialize new shell
    sh = shell.NewShell("localhost:5001")
}
