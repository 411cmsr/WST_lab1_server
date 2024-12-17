package models



type SOAPFault struct {
	Envelope struct {
		Body struct {
			Fault struct {
				Code    string `xml:"faultcode"`
				Message string `xml:"faultstring"`
				Detail  struct {
					ErrorCode    string `xml:"errorCode"`
					ErrorMessage string `xml:"errorMessage"`
				} `xml:"detail"`
			} `xml:"Fault"`
		} `xml:"Body"`
	} `xml:"Envelope"`
}

const (
	ErrorRecordNotFoundCode    = "404"
	ErrorRecordNotFoundMessage = "Запись не найдена"
	ErrorRecordNotFoundDetail  = "Запрашиваемая запись отсутствует в базе данных."
)