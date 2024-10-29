package handlers

import (
	"WST_lab1_server/models"

	"fmt"
)

// Обработчик методов soap
func ProcessRequest(request models.Request) (models.Response, error) {
	var response models.Response

	switch request.Method {
	case "Add":
		result, err := models.Database.AddData(request.Data)
		if err != nil {
			response.Status = "Error"
			return response, err
		}
		response.Status = "Success"
		response.Result = result

	case "Update":
		result, err := models.Database.UpdateData(request.Data)
		if err != nil {
			response.Status = "Error"
			return response, err
		}
		response.Status = "Success"
		response.Result = result

	case "Get":
		result, err := models.Database.GetData(request.Data)
		if err != nil {
			response.Status = "Error"
			return response, err
		}
		response.Status = "Success"
		response.Result = result

	case "Delete":
		result, err := models.Database.DeleteData(request.Data)
		if err != nil {
			response.Status = "Error"
			return response, err
		}
		response.Status = "Success"
		response.Result = result

	case "Search":
		results, err := models.Database.SearchData(request.Data)
		if err != nil {
			response.Status = "Error"
			return response, err
		}
		response.Status = "Success"
		response.Results = results

	case "GetAll":
		results, err := models.Database.GetAllData()
		if err != nil {
			response.Status = "Error"
			return response, err
		}
		response.Status = "Success"
		response.Results = results

	default:
		response.Status = "Error"
		return response, fmt.Errorf("unknown method: %s", request.Method)
	}

	return response, nil
}
