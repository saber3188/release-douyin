package main

import (
	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

func Init() {
	config.InitConfig()
}
func main() {
	Init()
	go service.RunMessageServer()
	r := gin.Default()

	initRouter(r)

	r.Run("192.168.1.27:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
