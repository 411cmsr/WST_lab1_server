package services

import (
	"WST_lab1_server/models"
	"encoding/xml"
)

type SearchRequest struct {
	XMLName xml.Name `xml:"searchRequest"`
	Query   string   `xml:"query"`
}

type SearchResponse struct {
	//XMLName xml.Name         `xml:"SearchResponse"`
	Results []models.Persons `xml:"SearchResult>SearchResult"`
}

//func SearchPersons(name, surname string, age *int) ([]models.Person, error) {
//	var persons []models.Person
//	query := config.DB.Model(&models.Person{})
//
//	if name != "" {
//		query = query.Where("name = ?", name)
//	}
//	if surname != "" {
//		query = query.Where("surname = ?", surname)
//	}
//	if age != nil {
//		query = query.Where("age = ?", *age)
//	}
//
//	err := query.Find(&persons).Error
//	return persons, err
//}
