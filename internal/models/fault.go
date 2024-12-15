package models

// import (
// 	"encoding/xml"
// )

type SOAPFault struct {
	//XMLName xml.Name `xml:"Fault"`
	FaultCode   string `xml:"faultcode"`
	FaultString string `xml:"faultstring"`
	FaultActor   string   `xml:"actor"`
	Detail      string `xml:"detail"`
}

// func (s *SOAPFault) Error() string {
// 	panic("unimplemented")
// }