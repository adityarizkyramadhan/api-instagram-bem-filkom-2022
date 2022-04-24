package handler

import (
	"api-instagram-bem-filkom-2022/helper"
	"api-instagram-bem-filkom-2022/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type idUser struct {
// 	Id string `uri :"id"`
// }

func UpdateDataInstagram(c *gin.Context) {
	err := service.GetResponseFromIG()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.ResponseAPI("APInya error kontak yang bikin ya", false, err))
		return
	}
	c.JSON(http.StatusOK, helper.ResponseAPI("Data diupdate", true, nil))
}

func GetDataFromDataBase(c *gin.Context) {
	data, err := service.GetData()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, helper.ResponseAPI("APInya error kontak yang bikin ya", false, err))
		return
	}
	c.JSON(http.StatusOK, helper.ResponseAPI("Data dari database berhasil ditemukan", true, data))
}
