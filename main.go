package main

import (
	"latihan-gin/src/config"
	"latihan-gin/src/helpers"
	"latihan-gin/src/routes"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	config.InitDB()
	helpers.Migrate()
	defer config.DB.Close()
	routes.Router()
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// routes.Router()
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080
}