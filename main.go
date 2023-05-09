package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dryRun := flag.Bool("d", false, "Dry run: print actions without actually deleting files")
	flag.Parse()

	dir := "."
	args := flag.Args()
	if len(args) > 0 {
		dir = args[0]
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	movFiles := make(map[string]string)

	for _, file := range files {
		if strings.ToLower(filepath.Ext(file.Name())) == ".mov" {
			nameWithoutExt := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
			movFiles[nameWithoutExt] = file.Name()
		}
	}

	for _, file := range files {
		if strings.ToLower(filepath.Ext(file.Name())) == ".heic" {
			nameWithoutExt := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))

			if movFile, ok := movFiles[nameWithoutExt]; ok {
				if *dryRun {
					fmt.Printf("Dry run: would delete file '%s'\n", filepath.Join(dir, movFile))
				} else {
					err := os.Remove(filepath.Join(dir, movFile))
					if err != nil {
						fmt.Printf("Error deleting file '%s': %v\n", filepath.Join(dir, movFile), err)
					} else {
						fmt.Printf("Deleted file '%s'\n", filepath.Join(dir, movFile))
					}
				}
			}
		}
	}
}
