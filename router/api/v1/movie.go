package v1

import (
	"github.com/OliCoder/RecSys/e"
	"github.com/OliCoder/RecSys/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMovieLists(c *gin.Context) {
	var movieLists []models.Movie

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetErrorCodeMsg(code),
		"data": movieLists,
	})
}
