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
type FolderStructure struct {
	DirName string
	Files   []string
}

// parserFlag initializes and returns a Config struct with parsed flag values.

func parserFlag() Config {
	now := time.Now().UTC()
	var cfg Config

	flag.StringVar(&cfg.pname, "pname", now.Format("20060102")+"_NewProject", "Name of the Project/parent directory")
	flag.StringVar(&cfg.path, "path", "./", "Path of where the project is going to live.")
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
	folders := []FolderStructure{
		{
			DirName: cfg.pname,
			Files:   []string{"README.md", "LICENSE", ".gitignore", ".env.example"},
		},
		{
			DirName: "src",
			Files:   []string{"__init__.py", "main.py"},
		},
		{
			DirName: "modules",
			Files:   []string{"__init.py__", "scanner.py", "parser.py", "utils.py"},
		},
		{
			DirName: "reporting",
			Files:   []string{"__init__.py", "repoter.py"},
		},
		{
			DirName: "tests",
			Files:   []string{"test_scanner.py", "test_reporter.py"},
		},
		{
			DirName: "infra",
			Files:   []string{"provider.tf", "main.tf", "outputs.tf", "variables.tf", "terraform.tfvars", "README.md"},
		},
		{
			DirName: "ci",
			Files:   []string{"pipeline.yaml"},
		},
		{
			DirName: "scripts",
			Files:   []string{"deploy.sh", "install.sh", "cleanup.sh"},
		},
		{
			DirName: "configs",
			Files:   []string{"config.yaml", "aws_config.yaml"},
		},
		{
			DirName: "docs",
			Files:   []string{"architecture.md", "usage.md", "changelog.md"},
		},
	}
	for _, direct := range folders {

		if direct.DirName == cfg.pname {
			err := os.Mkdir(cfg.path+"/"+direct.DirName, 0755)
			if err != nil {
				fmt.Printf("Failed to create folder %v\n", err)
			}
			for _, file := range direct.Files {
				f, err := os.Create("./" + direct.DirName + "/" + file)
				if err != nil {
					fmt.Printf("Error creating file: %v\n", err)
					return
				}
				defer f.Close()
			}

		} else {
			err := os.MkdirAll(cfg.path+"/"+cfg.pname+"/"+direct.DirName, 0755)
			if err != nil {
				fmt.Printf("Failed to create folder %s: %v\n", direct, err)
			}
			for _, file := range direct.Files {
				fle, err := os.Create("./" + cfg.pname + "/" + direct.DirName + "/" + file)
				if err != nil {
					fmt.Printf("Error creating file: %v\n", err)
					return
				}
				defer fle.Close()
			}

		}
	}
}
