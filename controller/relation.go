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

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	relationReq := &RelationReq{}
	err := c.ShouldBind(relationReq)
	if err != nil {
		log.Errorf("bind err the err is %s", err)
		return
	}
	user, err := cache.GetTokenInfo(relationReq.Token)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: err.Error()})
		log.Errorf("cache err the err is %s", err)
		return
	}
	id, err := strconv.ParseInt(relationReq.ToUserID, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: err.Error()})
		log.Errorf("parse err the err is %s", err)
		return
	}
	followerID := user.Id
	relation := &model.Relation{
		ID:         id,
		FollowerID: followerID,
	}
	if relationReq.ActionType == "1" {
		err = dao.CreatRelation(relation)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		}
	} else if relationReq.ActionType == "2" {
		err := dao.DelRelation(relation)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, model.Response{StatusCode: 0, StatusMsg: "关注成功"})
	log.Info("Relation success")
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	userID := c.Query("user_id")
	uID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		log.Errorf("parse err the err is %s", err)
		c.JSON(http.StatusOK, UserListResponse{
			Response: model.Response{
				StatusCode: 0,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	followList, err := dao.GetFollowList(uID)
	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: model.Response{
				StatusCode: 0,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		UserList: followList,
	})
	log.Info("get followList success")
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		UserList: []model.User{DemoUser},
	})
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		UserList: []model.User{DemoUser},
	})
}
