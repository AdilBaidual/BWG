package handler

import (
	"BWG/entity"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) createAccountHandler(c *gin.Context) {
	var tmpAccount entity.Account
	if err := c.BindJSON(&tmpAccount); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		logrus.Error(err.Error())
		return
	}

	ok, err := checkCurrencyCode(tmpAccount.CurrencyCode)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logrus.Error(err.Error())
		return
	}

	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "invalid currency code")
		return
	}

	id, err := h.services.Account.CreateAccount(tmpAccount)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logrus.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, id)
}

func (h *Handler) getAccountsHandler(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid user id param")
		logrus.Error(err.Error())
		return
	}

	accounts, err := h.services.GetAccountsByUserId(userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		logrus.Error(err.Error())
		return
	}

	c.JSON(http.StatusOK, accounts)
}
