package main

import (
	"Douyin/config"
	"Douyin/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := Initial(); err != nil {
		panic(err)
	}
	r := gin.Default()
	initRouter(r)
	err := r.Run() // listen and serve on localhost:8080
	if err != nil {
		panic(err)
	}
}

func Initial() error {
	//go message.RunMessageServer()
	if err := config.Init(); err != nil {
		return err
	}
	if err := repository.Init(); err != nil {
		return err
	}

	return nil
}
