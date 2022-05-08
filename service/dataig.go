package service

import (
	"api-instagram-bem-filkom-2022/config"
	"api-instagram-bem-filkom-2022/model"
	"context"
)

var db, _ = config.InitDB()

func addData(data []model.DataIG) error {
	db_manual := config.GetDB()
	test, err := GetData()
	if err != nil {
		return err
	}
	if test != nil {
		db_manual.ExecContext(context.Background(), "TRUNCATE TABLE data_igs;")
	}
	if err := db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func GetData() ([]model.DataIG, error) {
	// Get data with last created
	var data []model.DataIG
	// get lastest data from created_at
	db.Find(&data)
	if len(data) == 0 {
		return nil, nil
	}
	return data, nil
}
