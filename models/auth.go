package models

import (
	log "github.com/sirupsen/logrus"
)

type Claim struct {
	UserId  int  `gorm:"column:userid" json:"userid"`
	Level   int  `gorm:"column:level" json:"level"`
	IsAdmin bool `gorm:"column:isadmin" json:"isadmin"`
}

type Auth struct {
	UserId   int    `gorm:"column:userid;primary_key" json:"userid"`
	UserName string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
}

type ExtraClaim struct {
	UserName string
	Claims   []Claim
}

func CheckAuth(userName, password string) bool {
	var auth Auth
	log.Infof("username:%s password:%s try to login", userName, password)
	db.Select("userid").Where(Auth{UserName: userName, Password: password}).First(&auth)
	if auth.UserId > 0 {
		return true
	}
	return false
}

func GetExtraClaims(userName string) (claims []Claim) {
	userId := GetUserId(userName)
	db.Where("userid = ?", userId).Find(&claims)
	return
}

func GetUserId(userName string) int {
	var auth Auth
	db.Where("username = ?", userName).First(&auth)
	return auth.UserId
}
