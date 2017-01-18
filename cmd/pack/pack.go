// The pack command generates a template package from a template folder.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// given the directory for a template
// set each HTML file content as a const
// create an array of template files
// set each CSS file content as a const
// create an array of css contents to export

func main() {
	var templateDirectory string

	if len(os.Args) > 1 {
		templateDirectory = os.Args[1]
	} else {
		templateDirectory = "."
	}

	files, err := ioutil.ReadDir(templateDirectory)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found files: ")
	for _, file := range files {
		fmt.Printf("%v ", file.Name())
	}
	fmt.Println()
}

func packageTemplate(templateDir, packageDir string) {
	// read in template files
	// generate go files into package directory
}

func generatePackage(in templateInfo, packageDir string) {

}

func getFiles(templateDir string) ([]os.FileInfo, error) {
	files := make([]os.FileInfo, 0)
	files, err := ioutil.ReadDir(templateDir)
	if err != nil {
		return nil, err
	}
	return files, nil
}
