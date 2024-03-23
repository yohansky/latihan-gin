package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CustomHeaderAPI(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")

	c.JSON(http.StatusOK, gin.H{
		"message": "this response has custom headers",
	})
} //mengatur sumber daya
