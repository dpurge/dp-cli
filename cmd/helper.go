package cmd

import (
	"io/ioutil"
	"path/filepath"
)

func getBookFiles(dirname string, extension string) ([]string, error) {
	items, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	pages := make([]string, 0, len(items))

	for _, item := range items {
		if item.Mode().IsRegular() && filepath.Ext(item.Name()) == extension {
			fullname, err := filepath.Abs(filepath.Join(dirname, item.Name()))
			if err != nil {
				return nil, err
			}
			pages = append(pages, fullname)
		}
	}

	return pages, nil
}
