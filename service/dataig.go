package service

import (
	"api-instagram-bem-filkom-2022/config"
	"api-instagram-bem-filkom-2022/model"
)

var db, _ = config.InitDB()

func addData(data []model.DataIG) error {
	if err := db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetData() ([]model.DataIG, error) {
	// Get data with last created
	var data []model.DataIG
	// get lastest data from created_at
	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
