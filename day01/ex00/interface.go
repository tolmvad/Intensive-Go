package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type DBReader interface {
	recipesParser() Recipes
	recipesPrinter(data Recipes)
}

type Xml struct{}

func (Xml) recipesParser() Recipes {
	xmlFile, err := os.Open(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer xmlFile.Close()
	var data Recipes
	byteValue, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		log.Fatal(err)
	}
	err = xml.Unmarshal(byteValue, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func (Xml) recipesPrinter(data Recipes) {
	jsonOut, err := json.MarshalIndent(&data, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonOut))
}

type Json struct{}

func (Json) recipesParser() Recipes {
	jsonFile, err := os.Open(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	var data Recipes
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func (Json) recipesPrinter(data Recipes) {
	xmlOut, err := xml.MarshalIndent(&data, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(xmlOut))
}
