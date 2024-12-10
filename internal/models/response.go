package models
type GetAllPersonsResponse struct {
	Persons []Person `xml:"person"`
}
type GetPersonResponse struct {
	Person Person `xml:"person"`
}