package main

import "encoding/xml"

type Ingredient struct {
	XMLName xml.Name `xml:"item,omitempty" json:"-"`
	Name    string   `xml:"itemname,omitempty" json:"ingredient_name,omitempty"`
	Count   string   `xml:"itemcount,omitempty" json:"ingredient_count,omitempty"`
	Unit    string   `xml:"itemunit,omitempty" json:"ingredient_unit,omitempty"`
}

type Cake struct {
	XMLName     xml.Name     `xml:"cake,omitempty" json:"-"`
	Name        string       `xml:"name,omitempty" json:"name,omitempty"`
	Time        string       `xml:"stovetime,omitempty" json:"time,omitempty"`
	Ingredients []Ingredient `xml:"ingredients>item,omitempty" json:"ingredients,omitempty"`
}

type Recipes struct {
	XMLName xml.Name `xml:"recipes,omitempty" json:"-"`
	Cakes   []Cake   `xml:"cake,omitempty" json:"cake,omitempty"`
}
