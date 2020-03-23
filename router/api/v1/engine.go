package v1

import (
	"github.com/OliCoder/RecSys/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetEngineInfo(c *gin.Context) {
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetErrorCodeMsg(code),
		"data": "engine info",
	})
}

func UpdateEngineInfo(c *gin.Context) {
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetErrorCodeMsg(code),
		"data": "Update engine info success!",
	})
}
