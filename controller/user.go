package controller

import (
	"github.com/RaymondCode/simple-demo/internal/cache"
	"github.com/RaymondCode/simple-demo/internal/dao"
	"github.com/RaymondCode/simple-demo/internal/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sync/atomic"
)

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password
	existedUser, err := dao.GetUserByName(username)
	if err != nil {
		log.Errorf("Register|%v", err)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "Error happened,please try later"},
		})
		return
	}
	if existedUser != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
		return
	} else {
		atomic.AddInt64(&userIdSequence, 1)
		newUser := model.User{
			Id:       userIdSequence,
			Name:     username,
			PassWord: password,
			Token:    token,
		}
		if err := dao.CreateUser(&newUser); err != nil {
			log.Errorf("register err ,the err is %s", err)
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: model.Response{StatusCode: 1, StatusMsg: "Error happened,please try later"},
			})
			return
		}
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    username + password,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password
	user, err := dao.GetUserByName(username)
	if err != nil {
		log.Errorf("Login:GetUserByName err,the err is %s", err)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	if user == nil {
		log.Errorf("user doesn't exist")
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
		return
	}
	if user.Token != token {
		log.Errorf("password error")
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "password error"},
		})
		return
	}
	log.Info("Login success")
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: model.Response{StatusCode: 0},
		UserId:   user.Id,
		Token:    token,
	})
	if err = cache.SetTokenInfo(user, token); err != nil {
		log.Errorf("setToken err ,the err is %s", err)
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	user, err := cache.GetTokenInfo(token)
	if err != nil {
		log.Errorf("GetToken err,The err is %s", err)
		c.JSON(http.StatusOK, UserResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	if user == nil {
		log.Errorf("User doesn't exist")
		c.JSON(http.StatusOK, UserResponse{
			Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
		return
	}
	log.Info("User Info success")
	c.JSON(http.StatusOK, UserResponse{
		Response: model.Response{StatusCode: 0},
		User:     *user,
	})
}
