package handler

import (
	"BWG/entity"
	"BWG/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) getTransactionByIdHandler(c *gin.Context) {
	transactionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid transaction id param")
		logrus.Error(err.Error())
		return
	}

	transaction, err := h.services.Transaction.GetTransaction(transactionId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logrus.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func (h *Handler) getTransactionsHandler(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	perPageStr := c.DefaultQuery("per_page", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	options := entity.Options{
		Page:    page,
		PerPage: perPage,
	}

	transactions, err := h.services.GetTransactions(options)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logrus.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (h *Handler) getAccountTransactionsHandler(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid account id param")
		logrus.Error(err.Error())
		return
	}

	transactions, err := h.services.GetTransactionsByAccountId(userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logrus.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (h *Handler) invoiceHandler(c *gin.Context) {
	var invoice entity.Invoice
	if err := c.BindJSON(&invoice); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		logrus.Error(err.Error())
		return
	}

	invoice.TransactionStatus = service.TRANSACTION_CREATED

	ok, err := checkCurrencyCode(invoice.CurrencyCode)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logrus.Error(err.Error())
		return
	}

	if !ok {
		invoice.TransactionStatus = service.TRANSACTION_ERROR
	}

	if invoice.Amount < 0 {
		invoice.TransactionStatus = service.TRANSACTION_ERROR
	}

	id, err := h.services.Invoice(invoice)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, service.TRANSACTION_ERROR)
		logrus.Error(err.Error())
		return
	}

	transaction, err := h.services.GetTransaction(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, service.TRANSACTION_ERROR)
		logrus.Error(err.Error())
		return
	}

	if transaction.TransactionStatus == service.TRANSACTION_ERROR {
		newErrorResponse(c, http.StatusBadRequest, service.TRANSACTION_ERROR)
		return
	}

	c.JSON(http.StatusOK, id)
}

func (h *Handler) withdrawHandler(c *gin.Context) {
	var withdraw entity.Withdraw
	if err := c.BindJSON(&withdraw); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		logrus.Error(err.Error())
		return
	}

	withdraw.TransactionStatus = service.TRANSACTION_CREATED

	ok, err := checkCurrencyCode(withdraw.CurrencyCode)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logrus.Error(err.Error())
		return
	}

	if !ok {
		withdraw.TransactionStatus = service.TRANSACTION_ERROR
	}

	if withdraw.Amount < 0 {
		withdraw.TransactionStatus = service.TRANSACTION_ERROR
	}

	id, err := h.services.Withdraw(withdraw)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, service.TRANSACTION_ERROR)
		logrus.Error(err.Error())
		return
	}

	transaction, err := h.services.GetTransaction(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, service.TRANSACTION_ERROR)
		logrus.Error(err.Error())
		return
	}

	if transaction.TransactionStatus == service.TRANSACTION_ERROR {
		newErrorResponse(c, http.StatusBadRequest, service.TRANSACTION_ERROR)
		return
	}

	c.JSON(http.StatusOK, id)
}

func (h *Handler) initTransactionHandler(c *gin.Context) {
	transactionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid account id param")
		logrus.Error(err.Error())
		return
	}
	status, err := h.services.InitTransaction(transactionId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, service.TRANSACTION_ERROR)
		logrus.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, status)
}
