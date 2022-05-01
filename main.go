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
	r.Use(cors.Default())
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": true,
			"pesan":  "Hello KBMFILKOM!",
		})
	})
	r.GET("/data", handler.GetDataFromDataBase)
	r.GET("/update", handler.UpdateDataInstagram)
	r.Run()

	// status := strings.Contains("RUMAH ADVOKASI | Launching Rumah Advokasi", "RUMAH ADVOKASI")
	// fmt.Println(status)
	// Parse time "created_at": 1649509730,in golang
	// var waktu int64 = 1649509730
	// var created_at time.Time = time.Unix(waktu, 0)
	// // Format "DD-MM-YYYY"
	// fmt.Println(created_at.Format("02-01-2006"))
	// if created_at.Weekday() == time.Saturday {
	// 	fmt.Println("Sabtu")
	// }
	// fmt.Println(created_at)

}
