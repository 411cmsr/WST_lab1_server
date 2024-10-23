package models

type Persons interface {
	getName()
	getSurname()
	getAge()
	setSurname()
	GetDetails()
	toString()
}
type Person struct {
	ID      uint   `json:"id" gorm:"primary_key"` //id bigserial NOT NULL, CONSTRAINT "Persons_pkey" PRIMARY KEY (id)
	Name    string `json:"title"`                 //name character varying(200),
	Surname string `json:"author"`                // surname character varying(200),
	Age     int    //age integer,

}

// A person method
func (p Person) getName() {
	return p.Name
}

// A person method
func (p Person) getSurname() {
	return p.Surname
}

func (p Person) getAge() {
	return p.Age

}
func (p Person) setName(name string) {
	p.Name = name
}
func (p Person) setSurname(surname string) {
	p.Surname = surname
}
func (p Person) toString() string {
	return "Person{" + "name=" + p.Name + ", surname=" + p.Surname + ", age = " + p.Age + '}'
}
}
