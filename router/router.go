package router

import (
	"github.com/ariwanss/CvBackendGo/controller"
	"github.com/ariwanss/CvBackendGo/middleware"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func SetupRouter() {
	Router = gin.Default()
	Router.Use(middleware.ErrorMiddleware)

	Router.POST("/user/register", controller.Register, middleware.AttachToken)
	Router.POST("/user/login", controller.Login, middleware.AttachToken)
	Router.PUT("/user/update", middleware.Authorize, controller.UpdateUser)
	Router.DELETE("/user/delete", middleware.Authorize, controller.DeleteUser)

	// Router.GET("/all-users", controller.GetAllUser)

	cvRouter := Router.Group("/cv")
	cvRouter.Use(middleware.Authorize)
	cvRouter.POST("/", controller.CreateCvItem)
	cvRouter.GET("/", controller.GetCvItems)
	cvRouter.PUT("/:id", controller.UpdateCvItem)
	cvRouter.DELETE("/:id", controller.DeleteCvItem)

	// Router.POST("/cv", middleware.Authorize, controller.CreateCvItem)
	// Router.GET("/cv", middleware.Authorize, controller.GetCvItems)
	// Router.PUT("/cv/:id", middleware.Authorize, controller.UpdateCvItem)
	// Router.DELETE("/cv/:id", middleware.Authorize, controller.DeleteCvItem)

	// Router.GET("/cv-all", controller.GetAllCvItems)
	// Router.DELETE("/cv-all", controller.DropCvItemsCol)
}
