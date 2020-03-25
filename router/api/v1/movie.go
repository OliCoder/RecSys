package v1

import (
	"encoding/json"
	"github.com/OliCoder/RecSys/e"
	"github.com/OliCoder/RecSys/models"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
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
	movieRecentlyLists = models.GetMovieRecently(userId, req.MovieRecentlyNum)

	var movieTopRankLists []models.Movie
	movieTopRankLists = models.GetMovieTopRank(userId, req.MovieTopRankNum)

	var movieRecommendLists []models.Movie
	movieRecommendLists = models.GetMovieRecommend(userId, req.MovieRecommendNum)

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
