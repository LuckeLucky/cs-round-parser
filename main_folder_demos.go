//go:build !single
// +build !single

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/LuckeLucky/cs-round-parser/analyser"

	"github.com/LuckeLucky/cs-round-parser/utils"
)

func readDemos() {
	err := filepath.Walk("demos/",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			if filepath.Ext(path) != ".dem" {
				fmt.Println("Ignoring file: " + path)
				return nil
			}

			f, err := os.Open(path)
			utils.CheckError(err)
			defer f.Close()

			fmt.Printf("Analyzing file: %s\n", f.Name())
			an := analyser.NewAnalyser(f)
			an.FirstParse()
			fmt.Printf("Finished file: %s\n\n", f.Name())
			f.Close()
			return nil
		})
	if err != nil {
		panic(err)
	}
	fmt.Scanf("oi")
}
