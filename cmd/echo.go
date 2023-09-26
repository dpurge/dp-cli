package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("echo called")
		var err = (*struct{})(nil)
		if err != nil {
			log.Fatalln("This is a sample log message!")
		}
		fmt.Println("After log entry.")
	},
}

func init() {
	rootCmd.AddCommand(echoCmd)
}
