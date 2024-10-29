package models

//type Persons interface {
//	getName()
//	getSurname()
//	getAge()
//	setSurname()
//	GetDetails()
//	toString()
//}

type Persons struct {
	ID      uint   `gorm:"primaryKey; not null" xml:"id"`
	Name    string `gorm:"type:varchar(200)" xml:"name"`
	Surname string `gorm:"type:varchar(200)" xml:"surname"`
	Age     int    `json:"age,omitempty" xml:"age"`
}
