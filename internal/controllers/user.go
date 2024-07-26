package controllers

import (
	"context"
	"log/slog"
	"net/http"
	"paymentweb/internal/adapter/storage/postgres/repository"
	"paymentweb/internal/configuration"
	"paymentweb/internal/controllers/request"
	"paymentweb/internal/controllers/response"
	"paymentweb/internal/core/ports/in"
	"paymentweb/internal/core/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	ctx     *context.Context
	rg      *gin.RouterGroup
	service in.UserService
}

func NewUserController(ctx *context.Context, db configuration.DBContext, rg *gin.RouterGroup) {
	repo := repository.NewUserRepository(db)
	cardRepo := repository.NewCardRepository(db)
	cardService := service.NewCardService(cardRepo)
	service := service.NewUserService(repo, cardService)
	UserController{ctx, rg, service}.InitRoutes()
}

func (c UserController) InitRoutes() {
	users := c.rg.Group("/users")
	users.GET("/", c.list)
	users.GET("/:id", c.find)
	users.POST("/", c.create)
}

func (uc UserController) create(c *gin.Context) {
	var createUserRequest request.CreateUserRequest
	if err := c.ShouldBindJSON(&createUserRequest); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	if err := uc.service.Save(uc.ctx, createUserRequest.NewUserDomain()); err != nil {
		c.Error(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.Status(http.StatusCreated)
}

func (uc *UserController) list(c *gin.Context) {
	users, err := uc.service.List(uc.ctx)
	if err != nil {
		slog.Warn("Erro a listar usuários " + err.Error())
	}
	c.JSON(http.StatusOK, response.NewUserListResponse(users))
}

func (uc *UserController) find(c *gin.Context) {
	c.Param("id")
}
