package main

type Payload int

const (
	PayloadJSON Payload = iota
	PayloadXML
)

func (p Payload) String() string {
	switch p {
	case PayloadJSON:
		return "JSON"
	case PayloadXML:
		return "XML"
	default:
		return "UNKNOWN"
	}
}
