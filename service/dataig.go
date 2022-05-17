package service

import (
	"api-instagram-bem-filkom-2022/config"
	"api-instagram-bem-filkom-2022/model"
	"context"
)

var db, _ = config.InitDB()

func addData(data []model.DataIG) error {
	db_manual := config.GetDB()
	defer db_manual.Close()
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

func addDataSjw(data []model.DataIGSjw) error {
	db_manual := config.GetDB()
	defer db_manual.Close()
	test, err := GetData()
	if err != nil {
		return err
	}
	if test != nil {
		db_manual.ExecContext(context.Background(), "TRUNCATE TABLE data_ig_sjws;")
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

func GetDataSjw() ([]model.DataIGSjw, error) {
	// Get data with last created
	var data []model.DataIGSjw
	// get lastest data from created_at
	db.Find(&data)
	if len(data) == 0 {
		return nil, nil
	}
	return data, nil
}
