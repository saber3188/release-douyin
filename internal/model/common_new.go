package model

import "time"

type Response struct {
	StatusCode int32  `json:"status_code" gorm:"-"`
	StatusMsg  string `json:"status_msg,omitempty" gorm:"-"`
}

type Video struct {
	Id            int64     `json:"id,omitempty" gorm:"column:id;primaryKey"`
	Author        User      `json:"author" gorm:"embedded"`
	PlayUrl       string    `json:"play_url" gorm:"column:play_url"`
	CoverUrl      string    `json:"cover_url,omitempty" gorm:"column:cover_url"`
	FavoriteCount int64     `json:"favorite_count,omitempty" gorm:"column:favorite_count"`
	CommentCount  int64     `json:"comment_count,omitempty" gorm:"column:comment_count"`
	IsFavorite    bool      `json:"is_favorite,omitempty" gorm:"column:is_favorite"`
	CreatedAt     time.Time `json:"-" gorm:"column:created_at;index"`
	UpdatedAt     time.Time `json:"-" gorm:"column:updated_at"`
}

type Comment struct {
	Id         int64     `json:"id,omitempty" gorm:"column:id;primaryKey"`
	User       User      `json:"user" gorm:"embedded"`
	Content    string    `json:"content,omitempty" gorm:"column:content"`
	CreateDate string    `json:"create_date,omitempty" gorm:"column:create_date"`
	CreatedAt  time.Time `json:"-" gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"-" gorm:"column:updated_at"`
}

type User struct {
	Id            int64     `json:"id,omitempty" gorm:"column:id;primaryKey"`
	Name          string    `json:"name,omitempty" gorm:"column:name;index"`
	PassWord      string    `json:"pass_word" gorm:"column:pass_word"`
	Token         string    `json:"token" gorm:"column:token"`
	FollowCount   int64     `json:"follow_count,omitempty" gorm:"column:follow_count"`
	FollowerCount int64     `json:"follower_count,omitempty" gorm:"column:follower_count"`
	IsFollow      bool      `json:"is_follow,omitempty" gorm:"column:is_follow"`
	CreatedAt     time.Time `json:"-" gorm:"column:created_at;index"`
	UpdatedAt     time.Time `json:"-" gorm:"column:updated_at"`
}

type Message struct {
	Id         int64     `json:"id,omitempty" gorm:"column:id;primaryKey"`
	Content    string    `json:"content,omitempty" gorm:"column:content"`
	CreateTime string    `json:"create_time,omitempty" gorm:"column:create_time"`
	CreatedAt  time.Time `json:"-" gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"-" gorm:"column:updated_at"`
}

type MessageSendEvent struct {
	UserId     int64     `json:"user_id,omitempty" gorm:"column:user_id"`
	ToUserId   int64     `json:"to_user_id,omitempty" gorm:"column:to_user_id"`
	MsgContent string    `json:"msg_content,omitempty" gorm:"column:msg_content"`
	CreatedAt  time.Time `json:"-" gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"-" gorm:"column:updated_at"`
}

type MessagePushEvent struct {
	FromUserId int64     `json:"user_id,omitempty" gorm:"column:user_id"`
	MsgContent string    `json:"msg_content,omitempty" gorm:"column:msg_content"`
	CreatedAt  time.Time `json:"-" gorm:"column:created_at"`
	UpdatedAt  time.Time `json:"-" gorm:"column:updated_at"`
}
