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

type SearchParams struct {
	Name    string `xml:"name,omitempty"`
	Surname string `xml:"surname,omitempty"`
	Age     *int   `xml:"age,omitempty"` // Используем указатель для возможности передачи null
}

type SearchResponse struct {
	Persons []Persons `xml:"Persons>Person"`
}

//// A person method
//func (p Person) getName() {
//	return p.Name
//}

//func (p Person) getSurname() {
//	return p.Surname
//}

//func (p Person) getAge() {
//	return p.Age
//
//}
//func (p Person) setName(name string) {
//	p.Name = name
//}
//func (p Person) setSurname(surname string) {
//	p.Surname = surname
//}
//func (p Person) toString() string {
//	return "Person{" + "name=" + p.Name + ", surname=" + p.Surname + ", age = " + p.Age + '}'
//}
//}
