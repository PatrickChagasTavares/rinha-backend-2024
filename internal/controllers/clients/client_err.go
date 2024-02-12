package clients

import (
	"net/http"

	"github.com/patrickchagastavares/rinha-backend-2024/internal/services/clients"
	"github.com/patrickchagastavares/rinha-backend-2024/pkg/httpRouter"
)

func checkErr(c httpRouter.Context, err error) {
	switch err {
	case clients.ErrClientNotFound:
		c.JSON(http.StatusNotFound, "cliente não encontrado")
	case clients.ErrClientNotLimit:
		c.JSON(http.StatusUnprocessableEntity, "saldo insuficiente")
	default:
		c.JSON(http.StatusInternalServerError, "falha interna")
	}
}
