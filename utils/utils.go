package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func Md5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	str := hex.EncodeToString(h.Sum(nil))
	return str
}

//func GenerateSession(uname string) string {
//    return Md5String(fmt.Sprintf("%s:%d", uname, rand.Intn(999999)))
//}

func GenerateSession(userName string) string {
	return Md5String(fmt.Sprintf("%s:%s", userName, "session"))
}
func StringToTime(timeObj string) (time.Time, error) {
	timer, err := strconv.ParseInt(timeObj, 10, 64)
	log.Info("the timer is ", timer)
	if timer == 0 {
		return time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	}
	if err != nil {
		log.Info("parse err,the err is", err)
		return time.Time{}, err
	}
	if timer > 16905395820 {
		timer = timer / 1000
	}
	timeUnix := time.Unix(timer, 0).Format("2006-01-02 15:04:05")
	return time.Parse("2006-01-02 15:04:05", timeUnix)

}
