package dao

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/internal/model"
	"github.com/RaymondCode/simple-demo/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// GetUserByName 根据姓名获取用户
func GetUserByName(name string) (*model.User, error) {
	user := &model.User{}
	sqlDB := utils.GetDB()
	if sqlDB == nil {
		log.Info("get DB connect fail")
		return nil, fmt.Errorf("get DB connect fail")
	}
	if err := sqlDB.Model(model.User{}).Where("name=?", name).First(user).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, nil
		}
		log.Errorf("GetUserByName fail:%v", err)
		return nil, fmt.Errorf("GetUserByName fail:%v", err)
	}
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
