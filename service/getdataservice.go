package service

import (
	"api-instagram-bem-filkom-2022/model"

	"gorm.io/gorm"
)

func GetDataBemFilkom(db *gorm.DB) ([]model.DataIG, error) {
	var data []model.DataIG
	if err := db.Find(&data).Error; err != nil {
		return []model.DataIG{}, err
	}
	return data, nil
}

func GetDataSjwFilkom(db *gorm.DB) ([]model.DataIGSjw, error) {
	var data []model.DataIGSjw
	if err := db.Find(&data).Error; err != nil {
		return []model.DataIGSjw{}, err
	}
	return data, nil
}
