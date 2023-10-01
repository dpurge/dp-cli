package scan

import (
	"dpcli/cfg"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func GetPages(directory string, extension string) ([]string, error) {
	items, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	pages := make([]string, 0, len(items))

	for _, item := range items {
		if !item.IsDir() && filepath.Ext(item.Name()) == extension {
			fullname, err := filepath.Abs(filepath.Join(directory, item.Name()))
			if err != nil {
				return nil, err
			}
			pages = append(pages, fullname)
		}
	}

	return pages, nil
}

func CreateBlankPage(size string) (string, error) {
	magickConvert, err := cfg.GetTool("ImageMagick", "convert")
	if err != nil {
		return "", err
	}

	tmpFile, err := os.CreateTemp(".", "*.png")
	if err != nil {
		return "", err
	}
	tmpFile.Close()

	blank, err := filepath.Abs(tmpFile.Name())
	if err != nil {
		return "", err
	}

	cmd := exec.Command(magickConvert, "-size", size, "canvas:white", blank)
	buf, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	if len(buf) > 0 {
		log.Println(string(buf[:]))
	}

	return blank, nil
}

func CreateSignature(name string, signature []string) (string, error) {
	lenSignature := len(signature)
	if lenSignature%4 != 0 {
		log.Fatal(fmt.Sprintf("Number of pages in the signature (%d) in not divisible by 4!", lenSignature))
	}

	printSignature := make([]string, 0, lenSignature)

	for i := 0; i < lenSignature/2; i += 2 {
		printSignature = append(printSignature, signature[lenSignature-i-1], signature[i], signature[i+1], signature[lenSignature-i-2])
	}

	filename, err := CreatePdf(name, printSignature)
	if err != nil {
		return "", err
	}

	return filename, nil
}

func CreatePdf(name string, pages []string) (string, error) {
	filename, err := filepath.Abs(name + ".pdf")
	if err != nil {
		return "", err
	}

	magickConvert, err := cfg.GetTool("ImageMagick", "convert")
	if err != nil {
		return "", err
	}

	cmd := exec.Command(magickConvert, append(pages, filename)...)
	buf, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	if len(buf) > 0 {
		log.Println(string(buf[:]))
	}

	return filename, nil
}
