/*
Copyright © 2022 Saka
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/saka-naname/nyarn/cats"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available cats",
	Long: `list command will show a list of cats you can print.`,
	Run: func(cmd *cobra.Command, args []string) {
		chs := cats.GetCatHouse()
		w := tabwriter.NewWriter(os.Stdout, 0, 4, 2, ' ', 0)

		fmt.Println("Available Cats:")
		for _, c := range chs {
			w.Write([]byte("  " + c.Use + "\t" + c.Name + "\t\n"))
		}
		w.Flush()

		fmt.Println("\nUse \"nyarn [name]\" for print a cat")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
