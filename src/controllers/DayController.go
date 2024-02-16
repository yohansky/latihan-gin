package controllers

import (
	"latihan-gin/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SelectAllDay(c *gin.Context) {
	res := models.SelectAllDay()
	c.JSON(200, gin.H{
		"Message": "Tampil Semua Berhasil",
		"data": res,
	})
}

func SelectDay(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	res := models.SelectDay(id)
	c.JSON(200, gin.H{
		"Message": "Tampil Berhasil",
		"data": res,
	})
}

func InsertDay(c *gin.Context) {
	var input models.Day
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.InsertDay(input.Name)
	c.JSON(200, gin.H{
		"message": "Insert Berhasil",
	})
}

func UpdateDay(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	var input models.Day
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.UpdateDay(id, input.Name)
	c.JSON(200, gin.H{
		"message": "Update Berhasil",
	})
}

func DeleteDay(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	models.DeleteDay(id)
	c.JSON(200, gin.H{
		"message": "Delete Berhasil",
	})
}




