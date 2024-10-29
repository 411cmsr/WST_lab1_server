package models

type Request struct {
	Method string `xml:"method"`
	Data   string `xml:"data"`
}
