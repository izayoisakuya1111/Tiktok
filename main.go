package main

import (
  "github.com/gin-gonic/gin"
  "github.com/izayoisakuya1111/douyin/service"
)

//go mod download github.com/izayoisakuya1111/douyin
//go run main.go router.go

func main() {
	go service.RunMessageServer()

	r := gin.Default()

	initRouter(r)

	r.Run() 
}
