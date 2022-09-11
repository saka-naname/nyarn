/*
Copyright © 2022 Saka
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/saka-naname/nyarn/cats"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nyarn",
	Short: "Print a pretty cat",
	Long: `Nyarn is a parody application of Yarn. *meow*
(but this is NOT related with Yarn)
This application is a tool to print a pretty cat into your console.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cats.InitCats(rootCmd, cats.GetCatHouse())
}