package main

import (
	"fmt"
)

func main() {
	paths := []string{
		"data/payment/cc",
		"data/payment/cvv",
	}

	payloads := []struct {
		Body string
		Kind Payload
	}{
		{
			Body: SampleJSON,
			Kind: PayloadJSON,
		},
		{
			Body: SampleXML,
			Kind: PayloadXML,
		},
	}

	for _, payload := range payloads {
		fmt.Printf("Extracting %s =========\n", payload.Kind)
		values, err := ExtractPathValues(payload.Body, payload.Kind, paths)
		if err != nil {
			panic(err)
		}
		for _, value := range values {
			fmt.Println(value)
		}
		fmt.Println()
	}
}
