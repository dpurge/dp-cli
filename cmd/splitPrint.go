package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var splitPrintCmd = &cobra.Command{
	Use:   "split-print",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("split-print called")
		pages, err := getBookFiles(".", ".md")
		if err != nil {
			log.Fatal(err)
		}

		for _, page := range pages {
			fmt.Println(page)
		}
	},
}

func init() {
	rootCmd.AddCommand(splitPrintCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// printSplitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printSplitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
