package v1

import (
	"encoding/json"
	"github.com/OliCoder/RecSys/e"
	"github.com/OliCoder/RecSys/models"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetUserInfo(c *gin.Context) {
	var claims map[string]interface{}
	data, _ := c.Get("JWT_PAYLOAD")
	tmp, _ := json.Marshal(data.(jwt.MapClaims))
	_ = json.Unmarshal(tmp, &claims)
	log.Infof("context:%#v claims:%#v", c, claims)
	userName := claims["UserName"].(string)
	user := models.GetUserInfo(userName)

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetErrorCodeMsg(code),
		"data": user,
	})
}

func UpdateUserInfo(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Errorf("Get Update user info json failed, err:%v", err)
	}
	models.UpdateUserInfo(user)
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetErrorCodeMsg(code),
		"data": "Update user info success",
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
