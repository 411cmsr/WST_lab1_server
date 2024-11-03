package services

import (
	"WST_lab1_server/internal/models"
)

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

type GetAllPersonsResponse struct {
	Persons []models.Person `xml:"person"`
}

type SearchPersonRequest struct {
	Query string `xml:"query"`
}

func AddPersonRequestFactory() interface{} {
	return &AddPersonRequest{}
}

func UpdatePersonRequestFactory() interface{} {
	return &UpdatePersonRequest{}
}

func DeletePersonRequestFactory() interface{} {
	return &DeletePersonRequest{}
}

func GetPersonRequestFactory() interface{} {
	return &GetPersonRequest{}
}

func GetAllPersonsRequestFactory() interface{} {
	return &GetAllPersonsResponse{}
}

func SearchPersonRequestFactory() interface{} {
	return &SearchPersonRequest{}
}
