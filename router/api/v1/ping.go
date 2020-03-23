package v1

import (
	"context"
	"github.com/OliCoder/RecSys/e"
	"github.com/OliCoder/RecSys/engine"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func PingHandler(c *gin.Context) {
	client := engine.NewClient()
	ctx := context.Background()
	data, err := client.Ping(ctx)
	if err != nil {
		log.Error("Rpc call failed, err: %v", err)
		code := e.ERROR_RPC_CALL
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetErrorCodeMsg(code),
			"data": data,
		})
		return
	}
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetErrorCodeMsg(code),
		"data": data,
	})
}
