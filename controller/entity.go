package controller

import (
	"github.com/RaymondCode/simple-demo/internal/model"
)

const (
	avatar          = "http://192.168.1.27:8080/static/20210708221438_a7bee.jpg"
	backGroundImage = "http://192.168.1.27:8080/static/20210708221438_a7bee.jpg"
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
type UserListResponse struct {
	model.Response
	UserList []model.User `json:"user_list"`
}
type VideoListResponse struct {
	model.Response
	VideoList []model.Video `json:"video_list"`
}

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]model.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	model.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	model.Response
	User model.User `json:"user"`
}

var tempChat = map[string][]model.Message{}

var messageIdSequence = int64(1)

type ChatResponse struct {
	model.Response
	MessageList []model.Message `json:"message_list"`
}
type VideoListReq struct {
	Token  string `json:"token"`
	UserID string `json:"user_id"`
}
type FavoriteReq struct {
	Token      string `form:"token"`
	VideoID    string `form:"video_id"`
	ActionType string `form:"action_type"`
}
