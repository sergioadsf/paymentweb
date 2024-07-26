package controllers

import (
	"log"
	"net/http"
	"paymentweb/internal/configuration"

	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	rg *gin.RouterGroup
}

func NewPaymentController(rg *gin.RouterGroup) {
	PaymentController{rg}.InitRoutes()
}

func (c PaymentController) InitRoutes() {
	payments := c.rg.Group("/payments")
	payments.GET("/:id", Get)
	payments.POST("/", Save)
}

func Save(c *gin.Context) {

}

type Payment struct {
	Id   int
	Nome string
}

type GetPaymentRequest struct {
	ID int64 `uri:"id" binding:"required,min=1" example:"1"`
}

func Get(c *gin.Context) {
	var req GetPaymentRequest
	if err := c.ShouldBindUri(&req); err != nil {
		log.Print("Aqui", err.Error())
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	db := configuration.New().Get("pagamento")
	defer db.Close()
	stmt, err := db.Prepare("select id, nome from client where id = $1")
	if err != nil {
		log.Println("Erro ao preparar consulta, ", err.Error())
	}
	id := c.Param("id")
	var p Payment
	err = stmt.QueryRow(id).Scan(&p.Id, &p.Nome)
	if err != nil {
		log.Print("Nenhum pagamento para ID(", id, ")")
		c.JSON(http.StatusNotFound, "Nenhum pagamento encontrado")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ID":   p.Id,
		"Name": p.Nome,
	})
}
