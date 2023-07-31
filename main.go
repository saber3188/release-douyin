package main

import (
	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func Init() {
	config.InitConfig()
}
func main() {
	Init()
	go service.RunMessageServer()
	r := gin.Default()
	r.Use(gin.Logger())
	initRouter(r)

	err := r.Run("192.168.1.27:8080")
	if err != nil {
		log.Errorf("the err is %s", err)
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
