package clients

import (
	"github.com/patrickchagastavares/rinha-backend-2024/internal/controllers"
	"github.com/patrickchagastavares/rinha-backend-2024/pkg/httpRouter"
)

func New(router httpRouter.Router, ctrl *controllers.Container) {

	router.Post("/clientes/:id/transacoes", ctrl.Clients.Create)
	router.Get("/clientes/:id/extrato", ctrl.Clients.FindExtract)

}
