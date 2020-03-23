package models

type Claim struct {
	UserId  int
	Level   int
	IsAdmin bool
}

type Auth struct {
	UserId   int
	UserName string
	Password string
}

type ExtraClaim struct {
	UserName string
	Claims   []Claim
}

func CheckAuth(userName, password string) bool {
	var auth Auth
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
