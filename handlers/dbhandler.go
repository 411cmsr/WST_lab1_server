package handlers

import (
	"WST_lab1_server/models"
	"fmt"
	"gorm.io/gorm"
)

type DB struct {
	conn *gorm.DB
}

func NewDB(conn *gorm.DB) *DB {
	return &DB{conn: conn}
}

func (db *DB) AddData(data string) (uint, error) {
	item := models.Item{Data: data}
	if err := db.conn.Create(&item).Error; err != nil {
		return 0, fmt.Errorf("failed to add data: %w", err)
	}
	return item.ID, nil
}

func (db *DB) Update(id uint, data string) (uint, error) {
	var item models.Item
	if err := db.conn.First(&item, id).Error; err != nil {
		return 0, fmt.Errorf("no item found with id: %d", id)
	}

	item.Data = data
	if err := db.conn.Save(&item).Error; err != nil {
		return 0, fmt.Errorf("failed to update data: %w", err)
	}

	return item.ID, nil
}

func (db *DB) Get(id uint) (string, error) {
	var item models.Item
	if err := db.conn.First(&item, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", fmt.Errorf("no item found with id: %d", id)
		}
		return "", fmt.Errorf("failed to get data: %w", err)
	}

	return item.Data, nil
}

func (db *DB) Delete(id uint) (uint, error) {
	result := db.conn.Delete(&models.Item{}, id)

	if result.Error != nil {
		return 0, fmt.Errorf("failed to delete item: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return 0, fmt.Errorf("no item found with id: %d", id)
	}

	return id, nil
}

func (db *DB) Search(query string) ([]models.Item, error) {
	var results []models.Item

	if err := db.conn.Where("data ILIKE ?", "%"+query+"%").Find(&results).Error; err != nil {
		return nil, fmt.Errorf("failed to search data: %w", err)
	}

	return results, nil
}

func (db *DB) GetAll() ([]models.Item, error) {
	var results []models.Item

	if err := db.conn.Find(&results).Error; err != nil {
		return nil, fmt.Errorf("failed to get all data: %w", err)
	}

	return results, nil
}
