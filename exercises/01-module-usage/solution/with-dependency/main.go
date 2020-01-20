package main

import (
	"fmt"
	"github.com/buger/jsonparser"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("../../data/person.json")
	if err != nil {
		panic(err)
	}

	value, _, _, err :=jsonparser.Get(data, "city")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(value))
}
