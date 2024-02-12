package clients

import (
	"net/http"
	"strconv"

	"github.com/patrickchagastavares/rinha-backend-2024/internal/entities"
	"github.com/patrickchagastavares/rinha-backend-2024/internal/services"
	"github.com/patrickchagastavares/rinha-backend-2024/pkg/httpRouter"
	"github.com/patrickchagastavares/rinha-backend-2024/pkg/logger"
)

type (
	IController interface {
		Create(c httpRouter.Context)
		FindExtract(c httpRouter.Context)
	}

	controllers struct {
		srv *services.Container
		log logger.Logger
	}
)

func New(srv *services.Container, log logger.Logger) IController {
	return &controllers{
		srv: srv,
		log: log,
	}
}

func (ctrl *controllers) Create(c httpRouter.Context) {
	var transaction entities.TransactionRequest

	clientID, err := strconv.Atoi(c.GetParam("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	if err := c.DecodeJSON(&transaction); err != nil {
		c.String(http.StatusUnprocessableEntity, entities.ErrDecode.Error())
		return
	}

	transaction.ClientID = clientID
	if err := c.Validate(&transaction); err != nil {
		c.String(http.StatusUnprocessableEntity, err.Error())
		return
	}

	data, err := ctrl.srv.Clients.CreateTransaction(c.Context(), &transaction)
	if err != nil {
		checkErr(c, err)
		return
	}

	c.JSON(http.StatusOK, data)
}

func (ctrl *controllers) FindExtract(c httpRouter.Context) {
	clientID, err := strconv.Atoi(c.GetParam("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	data, err := ctrl.srv.Clients.FindExtract(c.Context(), clientID)
	if err != nil {
		checkErr(c, err)
		return
	}

	c.JSON(http.StatusOK, data)
}
