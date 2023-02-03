package main

import (
	"Douyin/controller/comment"
	"Douyin/controller/favorite"
	"Douyin/controller/feed"
	"Douyin/controller/message"
	"Douyin/controller/publish"
	"Douyin/controller/relation"
	"Douyin/controller/user"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// data directory is used to serve static resources
	r.Static("/static", "./data")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", feed.Feed)
	apiRouter.GET("/user/", user.UserInfo)
	apiRouter.POST("/user/register/", user.Register)
	apiRouter.POST("/user/login/", user.Login)
	apiRouter.POST("/publish/action/", publish.Publish)
	apiRouter.GET("/publish/list/", publish.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", favorite.FavoriteAction)
	apiRouter.GET("/favorite/list/", favorite.FavoriteList)
	apiRouter.POST("/comment/action/", comment.CommentAction)
	apiRouter.GET("/comment/list/", comment.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", relation.RelationAction)
	apiRouter.GET("/relation/follow/list/", relation.FollowList)
	apiRouter.GET("/relation/follower/list/", relation.FollowerList)
	apiRouter.GET("/relation/friend/list/", relation.FriendList)
	apiRouter.GET("/message/chat/", message.MessageChat)
	apiRouter.POST("/message/action/", message.MessageAction)
}
