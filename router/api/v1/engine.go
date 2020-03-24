package v1

import (
	"github.com/OliCoder/RecSys/e"
	"github.com/OliCoder/RecSys/settings"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
)

var EngineConf string

type EngineGroupReq struct {
	Data string `json:"data"`
}

func GetEngineInfo(c *gin.Context) {
	if len(EngineConf) == 0 {
		file, err := os.OpenFile(settings.GetConfigInstance().EngineConf, os.O_RDONLY, 0755)
		defer file.Close()
		if err != nil {
			log.Errorf("Open engine conf file failed, err:%v", err)
		}
		byteValue, err := ioutil.ReadAll(file)
		if err != nil {
			log.Errorf("Read json conf to memory failed, err:%v", err)
		}
		EngineConf = string(byteValue)
	}
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetErrorCodeMsg(code),
		"data": EngineConf,
	})
}

func UpdateEngineInfo(c *gin.Context) {
	var req EngineGroupReq
	err := c.BindJSON(&req)
	if err != nil {
		log.Errorf("get engine info from body failed, err:%v", err)
	} else {
		log.Infof("Update engine group json: %v", req.Data)
		EngineConf = req.Data
		//client := engine.NewClient()
		//ctx := context.Background()
		//status, err := client.UpdateEngineGroup(ctx, req.Data)
	}

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetErrorCodeMsg(code),
		"data": "Update engine info success!",
	})
}
