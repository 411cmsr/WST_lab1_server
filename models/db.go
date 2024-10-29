package models

type Item struct {
	ID   uint   `gorm:"primaryKey"`
	Data string `gorm:"type:text"`
}

type Database interface {
	AddData(data string) (uint, error)
	UpdateData(id uint, data string) (uint, error)
	GetData(id uint) (string, error)
	DeleteData(id uint) (uint, error)
	SearchData(query string) ([]Item, error)
	GetAllData() ([]Item, error)
}
