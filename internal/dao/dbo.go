package dao

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/internal/model"
	"github.com/RaymondCode/simple-demo/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

// GetUserByName 根据姓名获取用户
func GetUserByName(name string) (*model.User, error) {
	user := &model.User{}
	sqlDB := utils.GetDB()
	if sqlDB == nil {
		log.Errorf("get DB connect fail")
		return nil, fmt.Errorf("get DB connect fail")
	}
	if err := sqlDB.Model(&model.User{}).Where("name=?", name).First(user).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		log.Errorf("GetUserByName fail:%v", err)
		return nil, fmt.Errorf("GetUserByName fail:%v", err)
	}
	log.Info("GetUser Success")
	return user, nil
}
func CreateUser(user *model.User) error {
	if err := utils.GetDB().Model(&model.User{}).Create(user).Error; err != nil {
		log.Errorf("CreateUser fail: %v", err)
		return fmt.Errorf("CreateUser fail: %v", err)
	}
	log.Infof("insert success")
	return nil
}
func UpLoadVideo(video *model.Video) error {
	if err := utils.GetDB().Model(&model.Video{}).Create(video).Error; err != nil {
		log.Errorf("CreatVideo err,the err is %s", err)
		return fmt.Errorf("CreatVideo err,the err is %s", err)
	}
	log.Infof("Upload success")
	return nil
}
func GetVediosByTime(lastTime time.Time) ([]model.Video, int64, error) {
	var videoList []model.Video
	var count int64
	log.Info("lastTime:", lastTime)
	if err := utils.GetDB().Model(&model.Video{}).Order("created_at DESC").Where("created_at<?", lastTime).Find(&videoList).Limit(30).Count(&count).Error; err != nil {
		log.Errorf("GetVedio err,the err is %s", err)
		return nil, 0, err
	}
	log.Info("the count is ", count)
	return videoList, count, nil
}
func GetVideoListByUserID(user_id int64) ([]model.Video, error) {
	var videoList []model.Video
	if err := utils.GetDB().Model(&model.Video{}).Where("JSON_EXTRACT(user, '$.id') = ?", user_id).Find(&videoList).Error; err != nil {
		log.Errorf("GetVideoListByUserID err,the err is %s", err)
		return nil, err
	}
	return videoList, nil
}
func UpdateUser(user *model.User) error {
	if err := utils.GetDB().Model(&model.User{}).Where("id=?", user.Id).Updates(*user).Error; err != nil {
		log.Errorf("update err ,the err is %s", err)
		return err
	}
	return nil
}
func CreatFavorite(userId, videoId int64) error {
	favorite := &model.Favorite{
		UserId:  userId,
		VideoID: videoId,
	}
	if err := utils.GetDB().Model(&model.Favorite{}).Create(favorite).Error; err != nil {
		log.Errorf("Creat err the err is %s", err)
		return err
	}
	return nil
}
func DelFavorite(userID, videoID int64) error {
	if err := utils.GetDB().Model(&model.Favorite{}).Where("user_id=? AND video_id=?", userID, videoID).Delete(&model.Favorite{}).Error; err != nil {
		log.Errorf("Del err the err is %s", err)
		return err
	}
	return nil
}
func GetFavoriteVideoListByUserID(userID int64) ([]model.Video, error) {
	subquery := utils.GetDB().Model(&model.Favorite{}).Select("video_id").Where("user_id = ?", userID)
	var favoriteVideoList []model.Video
	if err := utils.GetDB().Model(&model.Video{}).Where("id IN (?)", subquery).Find(&favoriteVideoList).Error; err != nil {
		log.Errorf("db err the err is %s", err)
		return nil, err
	}
	return favoriteVideoList, nil
}
func CreatComment(comment *model.Comment) error {
	if err := utils.GetDB().Model(&model.Comment{}).Create(comment).Error; err != nil {
		log.Errorf("db err the err is %s", err)
		return err
	}
	return nil
}
func DelComment(commrntID int64) error {
	if err := utils.GetDB().Model(&model.Comment{}).Where("id =?", commrntID).Delete(&model.Comment{}).Error; err != nil {
		log.Errorf("db err the err is %s", err)
		return err
	}
	return nil
}
func GetCommentList(videoID int64) ([]model.Comment, error) {
	var commentList []model.Comment
	if err := utils.GetDB().Model(&model.Comment{}).Order("created_at DESC").Where("video_id=?", videoID).Find(&commentList).Error; err != nil {
		log.Errorf("db err the err is %s", err)
		return nil, err
	}
	return commentList, nil
}
func CreatRelation(relation *model.Relation) error {
	if err := utils.GetDB().Model(&model.Relation{}).Create(relation).Error; err != nil {
		log.Errorf("db err the err is %s", err)
		return err
	}
	return nil
}
func DelRelation(relation *model.Relation) error {
	if err := utils.GetDB().Model(&model.Relation{}).Where("id = ? AND follower_id = ?", relation.ID, relation.FollowerID).Error; err != nil {
		log.Errorf("db err the err is %s", err)
		return err
	}
	return nil
}
func GetFollowerList(userID int64) ([]model.User, error) {
	var followerList []model.User
	subquery := utils.GetDB().Model(&model.Relation{}).Select("follower_id").Where("id=?", userID)
	if err := utils.GetDB().Model(&model.User{}).Where("id IN (?)", subquery).Find(&followerList).Error; err != nil {
		log.Errorf("db err the err is %s", err)
		return nil, err
	}
	return followerList, nil
}
func GetFollowList(userID int64) ([]model.User, error) {
	var attentionList []model.User
	subquery := utils.GetDB().Model(&model.Relation{}).Select("id").Where("follower_id=?", userID)
	if err := utils.GetDB().Model(&model.User{}).Where("id IN (?)", subquery).Find(&attentionList).Error; err != nil {
		log.Errorf("db err the err is %s", err)
		return nil, err
	}
	return attentionList, nil
}
