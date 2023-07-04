package handlers

import (
	"github.com/gin-gonic/gin"
	cfg "gitlab.geogracom.com/skdf/skdf-abac-go/configs/app"
	"gitlab.geogracom.com/skdf/skdf-abac-go/internal/services"
	"gitlab.geogracom.com/skdf/skdf-abac-go/pkg/logger"
	"net/http"
)

type Handlers struct {
	Handler *Handler
}

type Option struct {
	Conf    cfg.Config
	Logger  logger.Logger
	Service *services.Service
}

func NewRouter(option Option) *gin.Engine {
	gin.SetMode(gin.TestMode)

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	h := &Handlers{
		Handler: New(&HandlerConfig{
			Conf:    option.Conf,
			Logger:  option.Logger,
			Service: option.Service,
		}),
	}

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Pong!",
		})
	})

	userAccess := router.Group("/user-access")
	{
		userAccess.POST("/create", h.Handler.AddUserAccess)
		userAccess.POST("/check", h.Handler.CheckUserAccess)
		userAccess.POST("/delete", h.Handler.DeleteUserAccess)
		userAccess.POST("/update", h.Handler.UpdateUserAccess)
	}

	return router
}
