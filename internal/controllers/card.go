package controllers

import (
	"context"
	"log/slog"
	"net/http"
	"paymentweb/internal/adapter/storage/postgres/repository"
	"paymentweb/internal/configuration"
	"paymentweb/internal/controllers/response"
	"paymentweb/internal/core/ports/in"
	"paymentweb/internal/core/service"

	"github.com/gin-gonic/gin"
)

type CardController struct {
	ctx     *context.Context
	rg      *gin.RouterGroup
	service in.CardService
}

func NewCardController(ctx *context.Context, db configuration.DBContext, rg *gin.RouterGroup) {
	repo := repository.NewCardRepository(db)
	service := service.NewCardService(repo)
	CardController{ctx, rg, service}.InitRoutes()
}

func (c CardController) InitRoutes() {
	cards := c.rg.Group("/cards")
	cards.GET("/", c.list)
	// users.GET("/:id", c.find)
	// users.POST("/", c.create)
}

func (cc CardController) list(c *gin.Context) {
	cards, err := cc.service.List(cc.ctx)
	if err != nil {
		slog.Warn("Erro a listar usuários " + err.Error())
	}
	c.JSON(http.StatusOK, response.NewCardListResponse(cards))
}
