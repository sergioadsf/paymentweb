package routes

import (
	"context"
	"paymentweb/internal/configuration"
	"paymentweb/internal/controllers"

	"github.com/gin-gonic/gin"
)

func Init(ctx *context.Context, db configuration.DBContext, r *gin.Engine) {
	v1 := r.Group("/v1")
	// controllers.NewPaymentController(v1)
	controllers.NewCardController(ctx, db, v1)
	controllers.NewUserController(ctx, db, v1)
}
