package controller

import "github.com/RaymondCode/simple-demo/internal/model"

var DemoVideos = []model.Video{
	{
		Id:            1,
		User:          DemoUser,
		PlayUrl:       "http://192.168.1.27:8080/static/812fd091-e5fe-4f2e-a9a3-68bd8a5204c5.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

var DemoComments = []model.Comment{
	{
		ID:      1,
		User:    DemoUser,
		Content: "Test Comment",
	},
}

var DemoUser = model.User{
	Id:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
