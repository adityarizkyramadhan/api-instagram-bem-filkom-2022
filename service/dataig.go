package service

import (
	"api-instagram-bem-filkom-2022/config"
	"api-instagram-bem-filkom-2022/model"
	"fmt"
)

var db, _ = config.InitDB()

func addData(data []model.DataIG) error {
	fmt.Println("===============================================================")
	fmt.Println(data)
	dataDB, _ := GetData()
	if dataDB == nil {
		if err := db.Create(&data).Error; err != nil {
			return err
		}
	}
	if len(data) != 0 {
		// replace dataDb lama ke data yang baru
		for _, dataNew := range data {
			for i, dataOld := range dataDB {
				if dataOld.LinkMedia == dataNew.LinkMedia {
					// update in database
					dataDB[i] = dataNew
				}
			}
		}
		// save data db yang lain ke array yang baru
		for _, dataNew := range data {
			for _, dataOld := range dataDB {
				if dataOld.LinkMedia != dataNew.LinkMedia {
					dataDB = append(dataDB, dataNew)
				}
			}
		}
	}
	for _, dataCreate := range dataDB {
		if err := db.Create(&dataCreate).Error; err != nil {
			return err
		}
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
