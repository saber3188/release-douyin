package cache

import (
	"encoding/json"
	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/internal/model"
	"github.com/RaymondCode/simple-demo/utils"
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
