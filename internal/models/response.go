package models

type GetAllPersonsResponse struct {
	Persons []Person `xml:"person"`
}
type GetPersonResponse struct {
	Person Person `xml:"person"`
}

type SOAPFault struct {
	FaultCode   string `xml:"faultcode"`
	FaultString string `xml:"faultstring"`
	Detail      string `xml:"detail"`
}

// Error implements error.
func (s *SOAPFault) Error() string {
	panic("unimplemented")
}

type DeleteResponse struct {
	Status bool `xml:"status"`
}
