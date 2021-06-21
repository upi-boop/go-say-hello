package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
)

// ShowAccount godoc
// @Summary Show a account
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param id path int true "Account ID"
// @Success 200 {object} Account
// @Header 200 {string} string	"ok"
// @Failure 400,404 {object} string	"error"
// @Failure 500 {object} string	"error"
// @Failure default {object} string	"error"
// @Router /accounts/{id} [get]
func (c *Controller) ShowAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	account, err := AccountOne(aid)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, account)
}

// ListAccounts godoc
// @Summary List accounts
// @Description get accounts
// @Accept  json
// @Produce  json
// @Param q query string false "name search by q"
// @Success 200 {array} Account
// @Header 200 {string} string	"ok"
// @Failure 400,404 {object} string	"error"
// @Failure 500 {object} string	"error"
// @Failure default {object} string	"error"
// @Router /accounts [get]
func (c *Controller) ListAccounts(ctx *gin.Context) {
	q := ctx.Request.URL.Query().Get("q")
	accounts, err := AccountsAll(q)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}

type Account struct {
	ID        int      `json:"id" example:"1"`
	Name      string   `json:"name" example:"account name"`
	PhotoUrls []string `json:"photo_urls" example:"http://test/image/1.jpg,http://test/image/2.jpg"`
}
