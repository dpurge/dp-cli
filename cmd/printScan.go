package cmd

import (
	"fmt"
	"log"

	"dpcli/scan"

	"github.com/spf13/cobra"
)

var name string
var directory string
var extension string

var printScanCmd = &cobra.Command{
	Use:   "print-scan",
	Short: "Create printable PDF signatures from a directory of scanned pages",
	Long: `Read scanned book pages from a directory, group them in signatures,
reorder pages for printing and convert to PDF.

Default format for the scanned pages is PNG.

Example 1:

	dp print-scan --name=my-book --directory=./my-book
	
Example 2:

	dp print-scan --name=my-book --directory=./testdata --extension=.png`,
	Run: func(cmd *cobra.Command, args []string) {
		pages, err := scan.GetPages(directory, extension)
		if err != nil {
			log.Fatal(err)
		}

		lenSignature := 32 // 4 pages * 8 sheets
		lenPages := len(pages)
		blank := "blank.png"

		signatureNr := 0
		for i := 0; i < lenPages; i += lenSignature {
			signature := make([]string, 0, lenSignature)
			signatureNr += 1

			j := i + lenSignature
			if j > lenPages {
				j = lenPages
			}
			signature = pages[i:j]
			for len(signature) < lenSignature {
				signature = append(signature, blank)
			}

			signatureName := fmt.Sprintf("%s-%02d", name, signatureNr)
			signatureFile, err := scan.CreateSignature(signatureName, signature)
			if err != nil {
				log.Fatal(err)
			}

			log.Println(signatureFile)
		}

		// delete blank page
	},
}

func init() {
	rootCmd.AddCommand(printScanCmd)

	printScanCmd.Flags().StringVarP(&name, "name", "n", "", "set book name")
	printScanCmd.MarkFlagRequired("name")

	printScanCmd.Flags().StringVarP(&directory, "directory", "d", "", "set directory with scanned pages")
	printScanCmd.MarkFlagRequired("directory")

	printScanCmd.Flags().StringVarP(&extension, "extension", "e", ".png", "set extension for scanned pages")
}
