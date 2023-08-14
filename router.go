package main

import (
	"github.com/gin-gonic/gin"
  	"github.com/izayoisakuya1111/douyin/controller"
)

func initRouter(r *gin.Engine) {
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)

}
