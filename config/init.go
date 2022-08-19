package config

import (
	"api-instagram-bem-filkom-2022/model"
	"database/sql"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DNS = "user=postgres password=adityarizky1020 host=db.jgjyjvyldoamqndazixl.supabase.co TimeZone=Asia/Singapore port=5432 dbname=postgres "

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&model.DataIG{}, &model.DataIGSjw{}); err != nil {
		return nil, err
	}
	return db, err
}

func GetDB() *sql.DB {
	db_manual, err := sql.Open("postgres", DNS)
	if err != nil {
		panic(err)
	}
	return db_manual
}
