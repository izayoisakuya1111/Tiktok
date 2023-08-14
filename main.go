package main

import (
  "github.com/gin-gonic/gin"
)

func main() {
	go RunMessageServer()

	r := gin.Default()

	initRouter(r)

	r.Run() 
}
