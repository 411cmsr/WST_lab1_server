package models

type AddPersonRequest struct {
	Name    string `xml:"name"`
	Surname string `xml:"surname"`
	Age     int    `xml:"age"`
}

type UpdatePersonRequest struct {
	ID      uint   `xml:"id"`
	Name    string `xml:"name"`
	Surname string `xml:"surname"`
	Age     int    `xml:"age"`
}

type DeletePersonRequest struct {
	ID uint `xml:"id"`
}

type GetPersonRequest struct {
	ID uint `xml:"id"`
}

type SearchPersonRequest struct {
	Query string `xml:"query"`
}
