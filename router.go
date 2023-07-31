package main

import (
	"github.com/RaymondCode/simple-demo/controller"
	"github.com/RaymondCode/simple-demo/internal/cache"
	"github.com/RaymondCode/simple-demo/internal/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")
	//basicRoute := r.Group("/douyin")
	apiRouter := r.Group("/douyin")
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.GET("/comment/list/", controller.CommentList)
	apiRouter.GET("/user/", controller.UserInfo)

	//apiRouter.Use(AuthMiddleWare())
	// basic apis
	apiRouter.POST("/publish/action/", controller.Publish)
	apiRouter.GET("/publish/list/", controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
	apiRouter.POST("/comment/action/", controller.CommentAction)

	// extra apis - II
	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
	apiRouter.GET("/relation/friend/list/", controller.FriendList)
	apiRouter.GET("/message/chat/", controller.MessageChat)
	apiRouter.POST("/message/action/", controller.MessageAction)
}
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if user, _ := cache.GetTokenInfo(token); user != nil && user.Token == token {
			log.Info("用户已验证")
			c.Next()
			return
		}
		log.Errorf("用户未验证")
		c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: 1,
			StatusMsg:  "用户未验证",
		})
		c.Abort()
		return
	}
}
