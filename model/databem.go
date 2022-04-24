package model

import "gorm.io/gorm"

type DataIG struct {
	*gorm.Model
	ThumbnailSrc string `gorm:"thumbnail_src"`
	Caption      string `gorm:"caption"`
	Tanggal      string `gorm:"tanggal"`
	Hari         string `gorm:"hari"`
	LinkMedia    string `gorm:"link_media"`
}
