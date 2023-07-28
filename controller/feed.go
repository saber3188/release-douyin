package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/internal/dao"
	"github.com/RaymondCode/simple-demo/internal/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type FeedRequest struct {
	latest_time string
	tokne       string
}
type FeedResponse struct {
	model.Response
	VideoList []model.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	feedReq := &FeedRequest{}
	err := c.ShouldBindQuery(feedReq)
	if len(feedReq.latest_time) == 0 {
		feedReq.latest_time = time.Now().Format("2006-01-02 15:04:05")
	}
	log.Info("the time is %s", feedReq.latest_time)
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		log.Errorf("Feed:bind err,the err is %s", err)
		return
	}
	lastTime, err := time.Parse("2006-01-02 15:04:05", feedReq.latest_time)
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
