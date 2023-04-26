package main

import (
	"fmt"
	"strings"

	"github.com/antchfx/jsonquery"
	"github.com/antchfx/xmlquery"
)

// ExtractPathValues extracts the values for a payload body based for each XPath given.
func ExtractPathValues(body string, kind Payload, paths []string) ([]any, error) {
	switch kind {
	case PayloadJSON:
		return extractPathValuesJSON(body, paths)
	case PayloadXML:
		return extractPathValuesXML(body, paths)
	}

	return nil, fmt.Errorf("unknown payload kind: %d", kind)
}

func extractPathValuesJSON(body string, paths []string) ([]any, error) {
	doc, err := jsonquery.Parse(strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	results := []any{}
	for _, path := range paths {
		node := jsonquery.FindOne(doc, path)
		if node == nil {
			return nil, fmt.Errorf("could not find value for path: %s", path)
		}
		results = append(results, node.Value())
	}
	return results, nil
}

func extractPathValuesXML(body string, paths []string) ([]any, error) {
	doc, err := xmlquery.Parse(strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	results := []any{}
	for _, path := range paths {
		node := xmlquery.FindOne(doc, path)
		if node == nil {
			return nil, fmt.Errorf("could not find value for path: %s", path)
		}
		results = append(results, node.InnerText())
	}
	return results, nil
}
