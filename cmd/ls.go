package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all available notes",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: parse database and list name plus hash
        fmt.Println("ls called")
    },
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
