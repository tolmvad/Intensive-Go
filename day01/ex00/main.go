package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func typeChoice() DBReader {
	var fileType DBReader
	var jsonType Json
	var xmlType Xml
	if strings.HasSuffix(os.Args[2], ".xml") {
		fileType = xmlType
	} else if strings.HasSuffix(os.Args[2], ".json") {
		fileType = jsonType
	} else {
		fmt.Fprintln(os.Stderr, "Error: wrong type of file, use *.xml or *.json type")
		os.Exit(1)
	}
	return fileType
}

func runPars() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Error: expect input", os.Args[0], "-f filename")
		os.Exit(1)
	}
	fileType := typeChoice()
	recipes := fileType.recipesParser()
	fileType.recipesPrinter(recipes)
}

func main() {
	flagF := flag.Bool("f", false, "use to input file")
	flag.Parse()
	if *flagF {
		runPars()
	} else {
		fmt.Fprintln(os.Stderr, "Error: expect input", os.Args[0], "-f filename")
		os.Exit(1)
	}
}
