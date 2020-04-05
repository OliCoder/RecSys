package models

import (
	"context"
	"github.com/OliCoder/RecSys/engine"
	log "github.com/sirupsen/logrus"
	"time"
)

type Movie struct {
	MovieId   int64   `json:"movieid" gorm:"column:movieid"`
	Title     string  `json:"title"`
	Genre     string  `json:"genre"`
	AvgRating float32 `json:"avgrating" gorm:"column:avgrating"`
	Img		  string  `json:"img"`
	Rating    int     `json:"rating"`
}

type Rating struct {
	Idx       int   `gorm:"column:idx"`
	UserId    int   `gorm:"column:userid"`
	MovieId   int   `gorm:"column:movieid"`
	Rating    int   `gorm:"column:rating"`
	TimeStamp int64 `gorm:"column:timestamp"`
}

func GetMovieRecently(userId, num int) []Movie {
	movies := make([]Movie, num)
	if num == 0 {
		return movies
	}
	var ratings []Rating
	db.Where("userid = ?", userId).Order("timestamp desc").Limit(num).Find(&ratings)
	for i := 0; i < len(ratings); i++ {
		db.Where("movieid = ?", ratings[i].MovieId).First(&movies[i])
		movies[i].Rating = ratings[i].Rating
	}
	return movies
}

func GetMovieTopRank(userId, num int) []Movie {
	movies := make([]Movie, num)
	if num == 0 {
		return movies
	}
	db.Order("avgrating desc").Limit(num).Find(&movies)
	for _, movie := range movies {
		var rating Rating
		if db.Select("userid = ? AND movieid = ?", userId, movie.MovieId).First(&rating).RecordNotFound() {
			continue
		}
		movie.Rating = rating.Rating
	}
	return movies
}

func GetMovieRecommend(userId, num int) []Movie {
	client := engine.NewClient()
	ctx := context.Background()
	timeThreshold := time.Now().Unix() - 60*60*24*7
	var cnt int
	db.Model(&Rating{}).Where("userid = ? AND timestamp >= ?", userId, timeThreshold).Count(&cnt)
	if cnt == 0 {
		cnt = 3
	}
	user := engine.UserProfile{
		UserId:                  int64(userId),
		MovieWatchedNumRecently: int32(cnt),
	}

	movieIdList, err := client.Recommend(ctx, &user, int32(num))
	log.Infof("Call rpc method Recommend user:%v num:%v", user, num)
	if err != nil {
		log.Errorf("Rpc call Recommend failed, err:%v", err)
	}
	movies := make([]Movie, len(movieIdList))
	for idx, id := range movieIdList {
		db.Where("movieid = ?", id).First(&movies[idx])
	}
	for _, movie := range movies {
		var rating Rating
		if db.Select("userid = ? AND movieid = ?", userId, movie.MovieId).First(&rating).RecordNotFound() {
			continue
		}
		movie.Rating = rating.Rating
	}
	return movies
}
