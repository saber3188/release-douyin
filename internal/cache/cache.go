package cache

import (
	"encoding/json"
	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/internal/model"
	"github.com/RaymondCode/simple-demo/utils"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"time"
)

func GetTokenInfo(token string) (*model.User, error) {
	redisKey := token
	val, err := utils.GetRedisCli().Get(context.Background(), redisKey).Result()
	if err != nil {
		return nil, err
	}
	user := &model.User{}
	err = json.Unmarshal([]byte(val), &user)
	return user, err
}
func SetTokenInfo(user *model.User, Token string) error {
	redisKey := Token
	val, err := json.Marshal(&user)
	if err != nil {
		return err
	}
	expired := time.Second * time.Duration(config.GetGlobalConf().Cache.SessionExpired)
	_, err = utils.GetRedisCli().Set(context.Background(), redisKey, val, expired*time.Second).Result()
	return err
}
func GetAllUser() ([]model.User, error) {
	keys, err := utils.GetRedisCli().Keys(context.Background(), "*").Result()
	if err != nil {
		log.Errorf("GetKeys fail,the err is %s", err)
		return nil, err
	}
	var value string
	var users []model.User
	var user model.User
	for _, key := range keys {
		value, err = utils.GetRedisCli().Get(context.Background(), key).Result()
		if err != nil {
			log.Errorf("GetValue err the err is %s", err)
			return nil, err
		}
		err = json.Unmarshal([]byte(value), &user)
		users = append(users, user)
	}
	return users, err
}
