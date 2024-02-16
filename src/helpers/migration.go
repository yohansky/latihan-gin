package helpers

import (
	"latihan-gin/src/config"
	"latihan-gin/src/models"
)

func Migrate() {
	config.DB.AutoMigrate(&models.Month{})
	config.DB.AutoMigrate(&models.Day{})
	config.DB.AutoMigrate(&models.Year{})
}