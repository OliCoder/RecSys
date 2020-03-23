package router

import (
	jwt "github.com/OliCoder/RecSys/middleware/app-jwt"
	"github.com/OliCoder/RecSys/middleware/cors"
	v1 "github.com/OliCoder/RecSys/router/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.CorsHandler())
	router.Use(gin.Recovery())
	var authMiddleware = jwt.GinJWTMiddlewareInit(jwt.AllAuthorizator)
	router.GET("/", v1.PingHandler)
	router.POST("/login", authMiddleware.LoginHandler)
	router.NoRoute(authMiddleware.MiddlewareFunc(), jwt.NoRouteHandler)
	auth := router.Group("/auth")
	{
		auth.GET("refresh_token", authMiddleware.RefreshHandler)
	}
	user := router.Group("/user")
	{
		user.GET("/info", v1.GetUserInfo)
		user.POST("/logout", v1.Logout)
	}
	api := router.Group("/api/v1")
	{
		api.GET("movie", v1.GetMovieLists)
	}
	admin := router.Group("/admin")
	var adminMiddleware = jwt.GinJWTMiddlewareInit(jwt.AdminAuthorizator)
	admin.Use(adminMiddleware.MiddlewareFunc())
	{
		admin.GET("/engine", v1.GetEngineInfo)
		admin.PUT("/engine", v1.UpdateEngineInfo)
	}
	return router
}
