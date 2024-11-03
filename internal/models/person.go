package models

type Person struct {
	ID      uint   `gorm:"primaryKey; not null" xml:"id" yaml:"id"`
	Name    string `gorm:"type:varchar(200)" xml:"name" yaml:"name"`
	Surname string `gorm:"type:varchar(200)" xml:"surname" yaml:"surname"`
	Age     int    `gorm:"age,omitempty" xml:"age" yaml:"age"`
}
