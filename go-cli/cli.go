package main

import (
	"fmt"
	"os"
)

var workingDir string

func init() {
	var err error
	workingDir, err = os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		os.Exit(1)
	}
}

func treeFolder() {

}

func main() {

	// err := os.Mkdir()
	fmt.Println(workingDir)

}
