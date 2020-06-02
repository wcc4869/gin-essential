package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wcc4869/ginessential/controller"
	"github.com/wcc4869/ginessential/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	r.GET("/", controller.Index)

	// 分类路由
	categoryRoutes := r.Group("/category")
	categoryRoutes.POST("/create", controller.CreateCategory)
	categoryRoutes.PUT("/update", controller.UpdateCategory)
	categoryRoutes.DELETE("/delete", controller.DeleteCategory)
	categoryRoutes.GET("/get", controller.GetCategory)
	categoryRoutes.GET("/list", controller.GetCategories)
	categoryRoutes.GET("/excel", controller.ExportExcel)

	return r
}
