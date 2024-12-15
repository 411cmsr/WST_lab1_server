package models

import (
	"encoding/xml"
)


type Envelope struct {
	XMLName xml.Name `xml:"http://www.w3.org/2003/05/soap-envelope Envelope"`
	Header  Header   `xml:"Header"`
	Body    Body     `xml:"Body"`
}


type Header struct {

}

type Body struct {
	AddPerson     *AddPersonRequest     `xml:"AddPerson,omitempty"`
	DeletePerson  *DeletePersonRequest  `xml:"DeletePerson,omitempty"`
	UpdatePerson  *UpdatePersonRequest  `xml:"UpdatePerson,omitempty"`
	GetPerson     *GetPersonRequest     `xml:"GetPerson,omitempty"`
	GetAllPersons *GetAllPersonsRequest `xml:"GetAllPersons,omitempty"`
	SearchPerson  *SearchPersonRequest  `xml:"SearchPerson,omitempty"`
	SOAPFault     *SOAPFault            `xml:"Fault,omitempty"`
}


