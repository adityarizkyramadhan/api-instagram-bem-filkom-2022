package main

import (
	"api-instagram-bem-filkom-2022/config"
	"api-instagram-bem-filkom-2022/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	_, err = config.InitDB()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	r.GET("/data", handler.GetDataBemFilkom)
	r.GET("/sjw", handler.GetDataBemFilkomSjw)
	r.GET("/data/update", handler.UpdateDataBemFilkom)
	r.GET("/sjw/update", handler.UpdateDataBemFilkomSjw)
	r.Run()

}
