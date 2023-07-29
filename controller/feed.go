package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/internal/dao"
	"github.com/RaymondCode/simple-demo/internal/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	feedReq := &FeedRequest{}
	feedReq.tokne = c.Query("token")
	feedReq.latest_time = c.Query("latest_time")
	var err error
	log.Info(feedReq)
	if len(feedReq.latest_time) == 0 {
		feedReq.latest_time = strconv.FormatInt(time.Now().Unix(), 10)
	}
	log.Info("the time is ", feedReq.latest_time)
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		log.Errorf("Feed:bind err,the err is %s", err)
		return
	}
	timeObj, _ := strconv.ParseInt(feedReq.latest_time, 10, 64)
	lastTimeUnix := time.Unix(timeObj, 0)
	lastTime, err := time.Parse("2006-01-02 15:04:05", lastTimeUnix.Format("2006-01-02 15:04:05"))
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		log.Errorf("parse err the err is %s", err)
	}
	VideoList, count, err := dao.GetVediosByTime(lastTime)
	if err != nil {
		log.Errorf("%s", err)
		c.JSON(http.StatusOK, FeedResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	if count == 0 {
		c.JSON(http.StatusOK, FeedResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "没有新视频哦"},
		})
		return
	}
	nextTime := VideoList[count-1].CreatedAt.Unix()
	c.JSON(http.StatusOK, FeedResponse{
		Response:  model.Response{StatusCode: 0},
		VideoList: VideoList,
		NextTime:  nextTime,
	})
	log.Info(fmt.Sprintf("the video list is %#v", VideoList))
	log.Info("Feed success")
}
