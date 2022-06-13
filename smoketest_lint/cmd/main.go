package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"github.com/greenbone/openvas-scanner/smoketest_lint/test"
)

func main() {

	openvasDir := flag.String("openvasDir", "", "location of the openvas-nasl-lint executable")
	path := flag.String("testFiles", "data", "folder for the nasl test data")
	flag.Parse()

	openvasExe := filepath.Join(*openvasDir, "openvas-nasl-lint")

	tfs := test.TestFiles()
	err := filepath.Walk(*path, tfs.Parse)
	if err != nil {
		fmt.Printf("Unable to parse files in %s: %s\n", *path, err)
	}

	for _, tf := range tfs.Tfs {
		errs := tf.Test(openvasExe)
		if len(errs) > 0 {
			fmt.Printf("%d error(s) while processing %s:\n", len(errs), tf.Name)
			for _, err := range errs {
				fmt.Println(err)
			}
		}
	}
}
