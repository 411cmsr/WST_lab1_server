package models

type Response struct {
	Status  string   `xml:"status"`
	Result  string   `xml:"result"`
	Results []string `xml:"results"`
}
