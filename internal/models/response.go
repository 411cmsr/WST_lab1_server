package models

type GetAllPersonsResponse struct {
	Persons []Person `xml:"person"`
}
type GetPersonResponse struct {
	Person Person `xml:"person"`
}



type DeleteResponse struct {
	Status bool `xml:"status"`
}

///////////////////////////////////////

type ErrorResponse struct {
    Type     string `json:"type"`
    Title    string `json:"title"`
    Status   int    `json:"status"`
    Detail   string `json:"detail"`
    Instance string `json:"instance"`
}