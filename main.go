package main

import (
  "github.com/gin-gonic/gin"
  "github.com/izayoisakuya1111/douyin/service"
)

func main() {
	go service.RunMessageServer()

	r := gin.Default()

	initRouter(r)

	r.Run() 
}
