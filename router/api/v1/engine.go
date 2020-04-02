package v1

import (
	"context"
	"fmt"
	"github.com/OliCoder/RecSys/e"
	"github.com/OliCoder/RecSys/engine"
	"github.com/OliCoder/RecSys/settings"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

var EngineConf string

type EngineGroupReq struct {
	Data interface{} `json:"data"`
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
	err := c.Bind(&req.Data)
	deleteSubStr := func(str, pattern string) string {
		reg, _ := regexp.Compile(pattern)
		return reg.ReplaceAllString(str, "")
	}
	if err != nil {
		fmt.Println(err)
		log.Errorf("get engine info from body failed, err:%v", err)
	} else {
		log.Infof("Update engine group json: %v", req)
		tmp := fmt.Sprintf("%#v", req.Data)
		tmp = deleteSubStr(tmp, `map\[string\]interface \{\}`)
		tmpByte := []byte(deleteSubStr(tmp, `\[\]interface \{\}`))
		tmpByte[10] = '['
		tmpByte[len(tmpByte)-2] = ']'

		EngineConf = string(tmpByte)
		client := engine.NewClient()
		ctx := context.Background()
		status, err := client.UpdateEngineGroup(ctx, EngineConf)
		if err != nil || status != true {
			log.Errorf("Rcp call UpdateEngineGroup failed, err:%v, status:%v", err, status)
		}
	}
	fmt.Println(EngineConf, req)
	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetErrorCodeMsg(code),
		"data": "Update engine info success!",
	})
}
