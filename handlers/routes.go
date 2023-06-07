package handlers

import (
	"ginCli/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *service.Service
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	apiRouter := router.Group("/tasks", h.AuthorizeMiddleware)
	apiRouter.GET("", h.GetAlLTasks)
	apiRouter.GET("/:id", h.GetTaskByID)
	apiRouter.GET("/by/:userID", h.GetTaskByIDUser)
	apiRouter.POST("", h.CreateTask)
	apiRouter.PUT("/:id", h.UpdateTask)
	apiRouter.DELETE("/:id", h.DeletedTask)

	userRouter := router.Group("/users", MiddlewareLogger, h.AuthorizeMiddleware)
	userRouter.GET("", h.GetAlLUsers)
	userRouter.GET("/by/:id", h.GetUserByID)
	userRouter.POST("", h.CreateUser)
	userRouter.PUT("/:id", h.UpdateUser)
	userRouter.DELETE("/:id", h.DeletedUser)

	authRouter := router.Group("/auth", MiddlewareLogger)
	authRouter.POST("sign-up", h.SignUp)
	authRouter.POST("sign-in", h.SignIn)

	return router
}
