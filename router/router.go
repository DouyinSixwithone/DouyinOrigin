package router

import (
	"Douyin/controller/comment"
	"Douyin/controller/favorite"
	"Douyin/controller/feed"
	"Douyin/controller/publish"
	"Douyin/controller/user"
	"Douyin/middleware/jwt"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	// data directory is used to serve static resources
	r.Static("/data", "./data")

	apiRouter := r.Group("/douyin")

	// basic apis 其中feed、register、login、publish/list接口对非登录用户也开放
	apiRouter.GET("/feed/", jwt.AuthWithoutLogin(), feed.List)
	apiRouter.GET("/user/", jwt.AuthWithLogin(), user.Info)
	apiRouter.POST("/user/register/", user.Register)
	apiRouter.POST("/user/login/", user.Login)
	apiRouter.POST("/publish/action/", jwt.AuthWithLogin(), publish.Action)
	apiRouter.GET("/publish/list/", jwt.AuthWithoutLogin(), publish.List)

	// extra apis - I 其中comment/list和favorite/list接口对非登录用户也开放
	apiRouter.POST("/favorite/action/", jwt.AuthWithLogin(), favorite.Action)
	apiRouter.GET("/favorite/list/", jwt.AuthWithoutLogin(), favorite.List)
	apiRouter.POST("/comment/action/", jwt.AuthWithLogin(), comment.Action)
	apiRouter.GET("/comment/list/", jwt.AuthWithoutLogin(), comment.List)

}
