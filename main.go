package main

import (
	"api-instagram-bem-filkom-2022/config"
	"api-instagram-bem-filkom-2022/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	_, err := config.InitDB()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	r.Use(cors.Default())
	r.GET("/data", handler.GetDataBemFilkom)
	r.GET("/sjw", handler.GetDataBemFilkomSjw)
	r.GET("/sjw/update", handler.UpdateDataBemFilkomSjw)
	r.Run()

}
