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
	comment := &model.Comment{
		VideoID: vID,
		User:    *user,
	}
	if actionType == "1" {
		comment.Content = content
		err = dao.CreatComment(comment)
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
	c.JSON(http.StatusOK, CommentActionResponse{Response: model.Response{StatusCode: 0},
		Comment: *comment,
	})
	log.Info("comment success")
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    model.Response{StatusCode: 0},
		CommentList: DemoComments,
	})
}
