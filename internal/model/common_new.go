package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Response struct {
	StatusCode int32  `json:"status_code" gorm:"-"`
	StatusMsg  string `json:"status_msg,omitempty" gorm:"-"`
}

type Video struct {
	Id int64 `json:"id,omitempty" gorm:"column:id;primaryKey"`
	//UserId        int64     `json:"user_id" gorm:"column:user_id"`
	User          User      `json:"user" gorm:"column:user"`
	Title         string    `json:"title" gorm:"column:title"`
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

func (u User) Value() (driver.Value, error) {
	// 这里我们将整个结构体转换成 JSON 格式存入数据库
	// 可根据需要进行其他格式的转换
	return json.Marshal(u)
}

// Scan 实现 sql.Scanner 接口，将数据库的值转换回结构体字段
func (u *User) Scan(value interface{}) error {
	// 这里我们从数据库中取出的值是 JSON 格式的数据
	// 将其解析并存入结构体字段
	if value == nil {
		return nil
	}
	return json.Unmarshal(value.([]byte), &u)
}
