package services

import (
	"WST_lab1_server/internal/models"
)


func AddPersonRequestFactory() interface{} {
	return &models.AddPersonRequest{}
}

func UpdatePersonRequestFactory() interface{} {
	return &models.UpdatePersonRequest{}
}

func DeletePersonRequestFactory() interface{} {
	return &models.DeletePersonRequest{}
}

func GetPersonRequestFactory() interface{} {
	return &models.GetPersonRequest{}
}

func GetAllPersonsRequestFactory() interface{} {
	return &models.GetAllPersonsResponse{}
}

func SearchPersonRequestFactory() interface{} {
	return &models.SearchPersonRequest{}
}
