package main

import (
    "fmt"
    "os"

    "github.com/ipfs/go-ipfs-api"

    "github.com/ex0dus-0x/jot/cmd"
)


// ipfs global bindings
var sh *shell.Shell
var ncalls int

func init() {
    fmt.Println("Doing init() stuff")
}

func main() {

    // perform argument parsing
    cmd.Execute()

    // initialize new shell
    sh = shell.NewShell("localhost:5001")

    os.Exit(0)
}
