package controller

import (
	"github.com/RaymondCode/simple-demo/internal/cache"
	"github.com/RaymondCode/simple-demo/internal/dao"
	"github.com/RaymondCode/simple-demo/internal/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	favoriteReq := &FavoriteReq{}
	err := c.ShouldBind(favoriteReq)
	log.Info("favorite req is ", favoriteReq)
	if err != nil {
		log.Errorf("Bind err the err is %s", err)
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	user, err := cache.GetTokenInfo(favoriteReq.Token)
	if err != nil {
		log.Errorf("GetToken err the err is %s", err)
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	videoID, err := strconv.ParseInt(favoriteReq.VideoID, 10, 64)
	if err != nil {
		log.Errorf("parse err the err is %s", err)
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	if favoriteReq.ActionType == "1" {
		user.FavoriteCount++
		err = cache.SetTokenInfo(user, favoriteReq.Token)
		if err != nil {
			log.Errorf("setCache err the err is %s", err)
			c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		}
		err = dao.CreatFavorite(user.Id, videoID)
		if err != nil {
			log.Errorf("db err the err is %s", err)
			c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		}
		c.JSON(http.StatusOK, model.Response{StatusCode: 0, StatusMsg: "已点赞"})
	} else if favoriteReq.ActionType == "2" {
		user.FavoriteCount--
		err = cache.SetTokenInfo(user, favoriteReq.Token)
		if err != nil {
			log.Errorf("setCache err the err is %s", err)
			c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		}
		err = dao.DelFavorite(user.Id, videoID)
		if err != nil {
			log.Errorf("db err the err is %s", err)
			c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		}
		c.JSON(http.StatusOK, model.Response{StatusCode: 0, StatusMsg: "已取消点赞"})
	}
	dao.UpdateUser(user)
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	favoriteListReq := &FavoriteListReq{}
	err := c.ShouldBind(favoriteListReq)
	if err != nil {
		log.Errorf("bind err the err is %s", err)
		c.JSON(http.StatusOK, VideoListResponse{
			Response: model.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	user, err := cache.GetTokenInfo(favoriteListReq.Token)
	if err != nil {
		log.Errorf("cache err ,the err is %s", err)
		c.JSON(http.StatusOK, VideoListResponse{
			Response: model.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	videoList, err := dao.GetFavoriteVideoListByUserID(user.Id)
	if err != nil {
		log.Errorf("the err is %s", err)
		c.JSON(http.StatusOK, VideoListResponse{
			Response: model.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		VideoList: videoList,
	})
	log.Info("FeedList success")
}
