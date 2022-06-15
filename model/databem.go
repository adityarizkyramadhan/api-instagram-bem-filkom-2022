package model

type DataIG struct {
	ThumbnailSrc string `gorm:"thumbnail_src"`
	Caption      string `gorm:"caption"`
	Tanggal      string `gorm:"tanggal"`
	Hari         string `gorm:"hari"`
	LinkMedia    string `gorm:"link_media"`
}

type DataIGSjw struct {
	ThumbnailSrc string `gorm:"thumbnail_src"`
	Caption      string `gorm:"caption"`
	Tanggal      string `gorm:"tanggal"`
	Hari         string `gorm:"hari"`
	LinkMedia    string `gorm:"link_media"`
}
