package main

import (
	"api-instagram-bem-filkom-2022/config"
	"api-instagram-bem-filkom-2022/handler"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
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
	store := persistence.NewInMemoryStore(5 * time.Hour)
	r.GET("/data", cache.CachePage(store, 5*time.Hour, handler.GetDataBemFilkom))
	r.GET("/sjw", cache.CachePage(store, 5*time.Hour, handler.GetDataBemFilkomSjw))
	r.Run()

}
