package app_jwt

import (
	"encoding/json"
	"github.com/OliCoder/RecSys/e"
	"github.com/OliCoder/RecSys/models"
	"github.com/OliCoder/RecSys/settings"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"time"
)

var indentity = settings.GetConfigInstance().IdentityKey

type JwtAuthorizator func(data interface{}, c *gin.Context) bool

func GinJWTMiddlewareInit(jwtAuthorizator JwtAuthorizator) (authMiddleware *jwt.GinJWTMiddleware) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "jwt middleware",
		SigningAlgorithm: "",
		Key:              []byte("secret key"),
		Timeout:          time.Hour * 5,
		MaxRefresh:       time.Hour,
		Authenticator: func(c *gin.Context) (interface{}, error) {
			log.Infof("gin context:%#v", c)
			var loginVal models.Auth
			if err := c.ShouldBind(&loginVal); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userName := loginVal.UserName
			password := loginVal.Password
			if models.CheckAuth(userName, password) {
				return &models.ExtraClaim{
					UserName: loginVal.UserName,
					Claims:   nil,
				}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: jwtAuthorizator,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			log.Infof("%#v", data)
			if v, ok := data.(*models.ExtraClaim); ok {
				v.Claims = models.GetExtraClaims(v.UserName)
				jsonClaims, _ := json.Marshal(v.Claims)
				return jwt.MapClaims{
					"UserName": v.UserName,
					"Claims":   string(jsonClaims),
				}
			}
			return jwt.MapClaims{}
		},
		Unauthorized: func(c *gin.Context, code int, msg string) {
			c.JSON(code, gin.H{
				"code": code,
				"msg":  msg,
			})
		},
		LoginResponse:   nil,
		LogoutResponse:  nil,
		RefreshResponse: nil,
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			jsonClaim := claims["Claims"].(string)
			var data []models.Claim
			json.Unmarshal([]byte(jsonClaim), &data)
			return &models.ExtraClaim{
				UserName: claims["UserName"].(string),
				Claims:   data,
			}
		},
		IdentityKey:           indentity,
		TokenLookup:           "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:         "Bearer",
		TimeFunc:              time.Now,
		HTTPStatusMessageFunc: nil,
		PrivKeyFile:           "",
		PubKeyFile:            "",
		SendCookie:            false,
		SecureCookie:          false,
		CookieHTTPOnly:        false,
		CookieDomain:          "",
		SendAuthorization:     false,
		DisabledAbort:         false,
		CookieName:            "",
	})

	if err != nil {
		log.Errorf("Jwt error, err:%v", err)
	}
	return
}

func AdminAuthorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(*models.ExtraClaim); ok {
		for _, item := range v.Claims {
			if item.IsAdmin && item.Level >= 3 {
				return true
			}
		}
	}
	return false
}

func AllAuthorizator(data interface{}, c *gin.Context) bool {
	return true
}

func NoRouteHandler(c *gin.Context) {
	code := e.PAGE_NOT_FOUND
	c.JSON(404, gin.H{
		"code": code,
		"msg":  e.GetErrorCodeMsg(code),
	})
}
