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

	item := router.Group("/item")
	{
		item.POST("/add", h.Handler.AddItem)
		item.POST("/delete", h.Handler.DeleteItem)
		item.POST("/update", h.Handler.UpdateItem)
	}

	router.POST("/table/edit", h.Handler.EditTable)

	return router
}
