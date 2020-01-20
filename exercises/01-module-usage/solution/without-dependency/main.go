package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func main() {
	byteValue, _ := ioutil.ReadFile("../../data/person.json")
	var person Person
	json.Unmarshal(byteValue, &person)
	fmt.Println(person)
}
