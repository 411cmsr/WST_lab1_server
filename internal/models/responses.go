package models

type GetAllPersonsResponse struct {
	Persons []Person `xml:"person"`
}
type SearchPersonResponse struct {
	Persons []Person `xml:"person"`
}
type GetPersonResponse struct {
	Person Person `xml:"person"`
}
type AddPersonResponse struct {
	ID uint `xml:"id"`
}
type UpdatePersonResponse struct {
	Success bool `xml:"success"`
}
type DeletePersonResponse struct {
	Success bool `xml:"success"`
}
