package v1

import (
	"encoding/json"
	"github.com/OliCoder/RecSys/e"
	"github.com/OliCoder/RecSys/models"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type GetMovieListReq struct {
	MovieRecentlyNum  int `json:"movie_recently_num"`
	MovieRecommendNum int `json:"movie_recommend_num"`
	MovieTopRankNum   int `json:"movie_top_rank_num"`
}

func GetMovieLists(c *gin.Context) {
	var claims map[string]interface{}
	data, _ := c.Get("JWT_PAYLOAD")
	tmp, _ := json.Marshal(data.(jwt.MapClaims))
	_ = json.Unmarshal(tmp, &claims)
	log.Infof("context:%#v claims:%#v", c, claims)
	userName := claims["UserName"].(string)
	userId := models.GetUserId(userName)

	var req GetMovieListReq
	err := c.BindJSON(&req)
	if err != nil {
		log.Errorf("Parse GetMovieListReq failed, err:%v, req:%v", err, req)
		code := e.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetErrorCodeMsg(code),
			"data": "failed to get movie lists",
		})
		return
	}

	var movieRecentlyLists []models.Movie
	if req.MovieRecentlyNum != 0 {
		movieRecentlyLists = models.GetMovieRecently(userId, req.MovieRecentlyNum)
	}

	var movieTopRankLists []models.Movie
	if req.MovieTopRankNum != 0 {
		movieTopRankLists = models.GetMovieTopRank(userId, req.MovieTopRankNum)
	}

	var movieRecommendLists []models.Movie
	if req.MovieRecommendNum != 0 {
		key := strconv.Itoa(userId) + "#" + strconv.Itoa(req.MovieRecommendNum)
		val, err := redisCli.Get(key).Result()
		if err == redis.Nil {
			movieRecommendLists = models.GetMovieRecommend(userId, req.MovieRecommendNum)
			go func() {
				val = ""
				for _, id := range movieRecommendLists {
					val += " " + strconv.FormatInt(id.MovieId, 10)
				}
				err := redisCli.Set(key, val[1:], time.Hour*24).Err()
				if err != nil {
					log.Errorf("Set redis key: %s failed, value: %s, err:%v", key, val[1:], err)
				} else {
					log.Infof("Set redis key: %s success, value: %s", key, val[1:])
				}
			}()
		} else if err != nil {
			log.Errorf("Get redis key: %v failed, err: %v", key, err)
		} else {
			log.Infof("Get redis key: %v success, val: %v", key, val)
			tmp := strings.Split(val, " ")
			var movieIdList []int64
			for _, item := range tmp {
				id, _ := strconv.ParseInt(item, 10, 64)
				movieIdList = append(movieIdList, id)
				movieRecommendLists = models.GetMovieListsInfo(userId, movieIdList)
			}
		}
	}

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetErrorCodeMsg(code),
		"data": gin.H{
			"movieRecently":  movieRecentlyLists,
			"movieTopRank":   movieTopRankLists,
			"movieRecommend": movieRecommendLists,
		},
	})
}
