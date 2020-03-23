package v1

import (
	"github.com/OliCoder/RecSys/e"
	"github.com/OliCoder/RecSys/models"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUserInfo(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userName := claims["UserName"].(string)
	user := models.GetUserInfo(userName)

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetErrorCodeMsg(code),
		"data": user,
	})
}

func Logout(c *gin.Context) {
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetErrorCodeMsg(code),
		"data": "logout success!",
	})
}
