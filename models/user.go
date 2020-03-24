package models

type User struct {
	Name       string
	UserId     int64 `gorm:"column:userid"`
	Age        int32
	Gender     string
	Occupation int
	ZipCode    int `gorm:"column:zipcode"`
}

var OccupationStatus = map[int]string{
	0:  "other",
	1:  "academic/educator",
	2:  "artist",
	3:  "clerical/admin",
	4:  "college/grad student",
	5:  "customer service",
	6:  "doctor/health care",
	7:  "executive/managerial",
	8:  "farmer",
	9:  "homemaker",
	10: "K-12 student",
	11: "lawyer",
	12: "programmer",
	13: "retired",
	14: "sales/marketing",
	15: "scientist",
	16: "self-employed",
	17: "technician/engineer",
	18: "tradesman/craftsman",
	19: "unemployed",
	20: "writer",
}

func GetOccupationStatus(code int) string {
	str, ok := OccupationStatus[code]
	if ok {
		return str
	}
	return OccupationStatus[0]
}

func GetUserInfo(userName string) (user User) {
	userId := GetUserId(userName)
	db.Where("userid = ?", userId).First(&user)
	return
}

func UpdateUserInfo(user User) {
	db.Save(&user)
}
