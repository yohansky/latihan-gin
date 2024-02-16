package controllers

import (
	"latihan-gin/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SelectAllYear(c *gin.Context) {
	res := models.SelectAllYear()
	c.JSON(200, gin.H{
		"Message": "Tampil Semua Berhasil",
		"data": res,
	})
}

func SelectYear(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	res := models.SelectYear(id)
	c.JSON(200, gin.H{
		"Message": "Tampil Berhasil",
		"data": res,
	})
}

func InsertYear(c *gin.Context) {
	var input models.Year
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.InsertYear(input.Name)
	c.JSON(200, gin.H{
		"message": "Insert Berhasil",
	})
}

func UpdateYear(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	var input models.Year
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.UpdateYear(id, input.Name)
	c.JSON(200, gin.H{
		"message": "Update Berhasil",
	})
}

func DeleteYear(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	models.DeleteYear(id)
	c.JSON(200, gin.H{
		"message": "Delete Berhasil",
	})
}
