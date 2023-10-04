package handler

import (
	"BWG/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoute() *gin.Engine {
	router := gin.New()

	api := router.Group("")
	{
		api.POST("/createUser", h.createUserHandler)
		api.GET("/getUser/:id", h.getUserByIdHandler)
		api.GET("/getUsers", h.getUsersHandler)

		api.POST("/createAccount", h.createAccountHandler)
		api.GET("/getUserAccounts/:id", h.getAccountsHandler)

		api.POST("/invoice", h.invoiceHandler)
		api.POST("/withdraw", h.withdrawHandler)
		api.POST("/initTransaction/:id", h.initTransactionHandler)
		api.GET("/getTransaction/:id", h.getTransactionByIdHandler)
		api.GET("/getTransactions", h.getTransactionsHandler)
		api.GET("/getAccountTransactions/:id", h.getAccountTransactionsHandler)
	}

	return router
}
