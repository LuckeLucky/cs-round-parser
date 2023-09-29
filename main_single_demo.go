//go:build single
// +build single

package main

import (
	"fmt"
	"os"

	"github.com/LuckeLucky/cs-round-parser/analyser"
	"github.com/LuckeLucky/cs-round-parser/utils"
)

func readDemos() {
	f, err := os.Open(os.Args[1])
	utils.CheckError(err)
	defer f.Close()

	fmt.Printf("Analyzing file: %s\n", f.Name())
	an := analyser.NewAnalyser(f)
	an.FirstParse()
	fmt.Printf("Finished file: %s\n\n", f.Name())
	f.Close()

	fmt.Scanf("oi")
}
