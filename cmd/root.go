package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// define our database file arg, if specified
var dbFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jot",
	Short: "IPFS-empowered decentralized dnote-like note-taking client;",
	Long: `jot is a lightweight note-taking app similar to dnote, except it stores
notes on the IPFS gateway`,
	Run: func(cmd *cobra.Command, args []string) { 
        fmt.Println("Printing pretty stuff!")
    },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	var err = rootCmd.Execute();
    if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initDB)

	// defines our global persistent flags
    rootCmd.PersistentFlags().StringVar(&dbFile, "database", "", "database file (default is $HOME/.note.jot)")

}

// initDB reads in database file for bolt consumption and ENV variables if set.
func initDB() {
	if dbFile != "" {
		// Use specified file from the flag.
		viper.SetConfigFile(dbFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".test" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".test")
	}

	viper.AutomaticEnv() // read in environment variables that match
}
