package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// Config holds the values parsed from command-line flags.
type Config struct {
	pname string
	path  string
}

// parserFlag initializes and returns a Config struct with parsed flag values.

func parserFlag() Config {
	now := time.Now().UTC()
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
	}
	var cfg Config

	flag.StringVar(&cfg.pname, "pname", now.Format("20060102")+"_NewProject", "Name of the Project/parent directory")
	flag.StringVar(&cfg.path, "path", cwd, "Path of where the project is going to live.")
	flag.Parse()

	return cfg
}

// main initializes configuration and passes it to other functions.

func main() {
	cfg := parserFlag()
	testfun(cfg)
}

// testfun uses the Config values to print project details.

func testfun(cfg Config) {
	directories := []string{"src", "tests", "infra", "ci", "scripts", "configs", "docs"}
	for _, direct := range directories {
		err := os.Mkdir(cfg.path+"\\"+direct, 0755)
		if err != nil {
			fmt.Printf("Failed to create folder %s: %v\n", direct, err)
		}
	}
}
