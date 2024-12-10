package models
type SearchPersonRequest struct {
	Query string `xml:"query"`
}

type AddPersonRequest struct {
	Name    string `xml:"name"`
	Surname string `xml:"surname"`
	Age     int    `xml:"age"`
	Email   string `xml:"email"`
	Telephone string `xml:"telephone"`
}

type UpdatePersonRequest struct {
	ID      uint   `xml:"id"`
	Name    string `xml:"name"`
	Surname string `xml:"surname"`
	Age     int    `xml:"age"`
	Email   string `xml:"email"`
	Telephone string `xml:"telephone"`
}

type DeletePersonRequest struct {
	ID uint `xml:"id"`
}

type GetPersonRequest struct {
	ID uint `xml:"id"`
}
