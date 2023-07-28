package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/internal/cache"
	"github.com/RaymondCode/simple-demo/internal/dao"
	"github.com/RaymondCode/simple-demo/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"net/http"
	"path/filepath"
)

const coverUrl = "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg"

type VideoListResponse struct {
	model.Response
	VideoList []model.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")
	user, _ := cache.GetTokenInfo(token)
	if user == nil {
		log.Errorf("Publish :user not exist")
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		log.Errorf("uploaded failed,the err is %s", err)
		return
	}
	log.Infof("upload success")
	FileName := uuid.New().String() + ".mp4"
	saveFile := filepath.Join("./public/", FileName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		log.Errorf("save failed,the err is %s", err)
		return
	}
	FinalName := fmt.Sprintf("http://192.168.1.27:8080/static/%s", FileName)
	video := &model.Video{
		User:     *user,
		Title:    title,
		PlayUrl:  FinalName,
		CoverUrl: coverUrl,
	}
	if err = dao.UpLoadVideo(video); err != nil {
		log.Errorf("upload err,the err is %s", err)
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		StatusCode: 0,
		StatusMsg:  FileName + " uploaded successfully",
	})
	log.Infof("save successfully ,the path is %s", FileName)
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
