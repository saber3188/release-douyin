package controller

import (
	"github.com/RaymondCode/simple-demo/internal/cache"
	"github.com/RaymondCode/simple-demo/internal/dao"
	"github.com/RaymondCode/simple-demo/internal/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

type CommentListResponse struct {
	model.Response
	CommentList []model.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	model.Response
	Comment model.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	videoID := c.Query("video_id")
	actionType := c.Query("action_type")
	commentID := c.Query("comment_id")
	content := c.Query("comment_text")
	user, _ := cache.GetTokenInfo(token)
	vID, err := strconv.ParseInt(videoID, 10, 64)
	if err != nil {
		log.Errorf("parse err.the err is %s", err)
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	Comment := &model.Comment{
		VideoID:   vID,
		User:      *user,
		CreatedAt: time.Now(),
	}
	if actionType == "1" {
		Comment.Content = content
		Comment.CreatDate = Comment.CreatedAt.Format("01-02")
		err = dao.CreatComment(Comment)
		if err != nil {
			log.Errorf("db err.the err is %s", err)
			c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		}
	} else if actionType == "2" {
		cID, err := strconv.ParseInt(commentID, 10, 64)
		if err != nil {
			log.Errorf("parse err.the err is %s", err)
			c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		}
		err = dao.DelComment(cID)
		if err != nil {
			log.Errorf("db err.the err is %s", err)
			c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		}
	}
	//c.JSON(http.StatusOK, model.Response{StatusCode: 0, StatusMsg: "没问题"})
	c.JSON(http.StatusOK, CommentActionResponse{Response: model.Response{StatusCode: 0, StatusMsg: "已评论"},
		Comment: *Comment,
	})
	//c.JSON(http.StatusOK, gin.H{
	//	"status_code": 0,
	//	"status_msg":  "meiwent",
	//	"comment":     *Comment,
	//})
	log.Info(" success")
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	videoID := c.Query("video_id")
	vid, err := strconv.ParseInt(videoID, 10, 64)
	if err != nil {
		log.Errorf("parse err the err is %s", err)
		c.JSON(http.StatusOK, CommentListResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	CommentList, err := dao.GetCommentList(vid)
	if err != nil {
		log.Errorf("db err the err is %s", err)
		c.JSON(http.StatusOK, CommentListResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	//c.JSON(http.StatusOK, model.Response{StatusCode: 0, StatusMsg: "测试"})
	c.JSON(http.StatusAccepted, CommentListResponse{
		Response:    model.Response{StatusCode: 0},
		CommentList: CommentList,
	})
	log.Info("GetCommentList success")
	log.Info(CommentList)
}
