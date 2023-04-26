package main

import (
	"fmt"
)

func main() {
	paths := []string{
		"data/payment/cc",
		"data/payment/cvv",
	}

	fmt.Println("JSON ===============")
	values, err := ExtractPathValues(SampleJSON, PayloadJSON, paths)
	if err != nil {
		panic(err)
	}
	for _, value := range values {
		fmt.Println(value)
	}

	fmt.Println()
	fmt.Println("XML ================")
	values, err = ExtractPathValues(SampleXML, PayloadXML, paths)
	if err != nil {
		panic(err)
	}
	for _, value := range values {
		fmt.Println(value)
	}
}
