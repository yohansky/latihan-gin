package routes

import (
	"latihan-gin/src/controllers"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()

	// // Simple group: v1
	// v1 := router.Group("/v1")
	// {
	// 	v1.POST("/login", loginEndpoint)
	// 	v1.POST("/submit", submitEndpoint)
	// 	v1.POST("/read", readEndpoint)
	// }

	// // Simple group: v2
	// v2 := router.Group("/v2")
	// {
	// 	v2.POST("/login", loginEndpoint)
	// 	v2.POST("/submit", submitEndpoint)
	// 	v2.POST("/read", readEndpoint)
	// }

	v1 := router.Group("/api/v1")
	{
		ping := v1.Group("/ping")
		{
			ping.GET("/", controllers.PingController)
		}
		month := v1.Group("/month")
		{
			month.GET("/list", controllers.MonthList)
			month.GET("/:id", controllers.MonthShow)
			month.POST("/create", controllers.MonthCreate)
			month.PUT("/update/:id", controllers.MonthUpdate)
			month.DELETE("/delete/:id", controllers.MonthDelete)
		}
		day := v1.Group("/day")
		{
			day.GET("/list", controllers.SelectAllDay)
			day.GET("/:id", controllers.SelectDay)
			day.POST("/create", controllers.InsertDay)
			day.PUT("/update/:id", controllers.UpdateDay)
			day.DELETE("/delete/:id", controllers.DeleteDay)
		}
		year := v1.Group("/year")
		{
			year.GET("/list", controllers.SelectAllYear)
			year.GET("/:id", controllers.SelectYear)
			year.POST("/create", controllers.InsertYear)
			year.PUT("/update/:id", controllers.UpdateYear)
			year.DELETE("/delete/:id", controllers.DeleteYear)
		}
	}
	router.Run(":8080")
}