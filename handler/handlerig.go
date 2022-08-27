package handler

import (
	"api-instagram-bem-filkom-2022/config"
	"api-instagram-bem-filkom-2022/helper"
	"api-instagram-bem-filkom-2022/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type idUser struct {
// 	Id string `uri :"id"`
// }

func GetDataBemFilkom(c *gin.Context) {
	db, err := config.InitDB()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.ResponseAPI("APInya error kontak yang bikin ya", false, err.Error()))
		return
	}
	data, err := service.GetDataBemFilkom(db)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.ResponseAPI("APInya error kontak yang bikin ya", false, err.Error()))
		return
	}
	c.JSON(http.StatusOK, helper.ResponseAPI("Data dari database berhasil ditemukan", true, data))
}

func GetDataBemFilkomSjw(c *gin.Context) {
	db, err := config.InitDB()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.ResponseAPI("APInya error kontak yang bikin ya", false, err.Error()))
		return
	}
	data, err := service.GetDataSjwFilkom(db)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.ResponseAPI("APInya error kontak yang bikin ya", false, err.Error()))
		return
	}
	c.JSON(http.StatusOK, helper.ResponseAPI("Data dari database berhasil ditemukan", true, data))
}

func UpdateDataBemFilkomSjw(c *gin.Context) {
	db, err := config.InitDB()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.ResponseAPI("APInya error kontak yang bikin ya", false, err.Error()))
		return
	}
	data, err := service.UpdateResponseFromHastag(db)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.ResponseAPI("APInya error kontak yang bikin ya", false, err.Error()))
		return
	}
	c.JSON(http.StatusOK, helper.ResponseAPI("Data dari database berhasil ditemukan", true, data))
}

// func GetDataSjw(c *gin.Context) {
// 	data, err := service.GetResponseFromHastag()
// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.ResponseAPI("APInya error kontak yang bikin ya", false, err.Error()))
// 		return
// 	}
// 	c.JSON(http.StatusOK, helper.ResponseAPI("Data dari database berhasil ditemukan", true, data))
// }
