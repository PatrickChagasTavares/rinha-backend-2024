package handlers

import (
	"os"

	"github.com/patrickchagastavares/rinha-backend-2024/internal/controllers"
	"github.com/patrickchagastavares/rinha-backend-2024/internal/handlers/clients"
	"github.com/patrickchagastavares/rinha-backend-2024/internal/handlers/swagger"
	"github.com/patrickchagastavares/rinha-backend-2024/pkg/httpRouter"
)

type (
	Options struct {
		Ctrl   *controllers.Container
		Router httpRouter.Router
	}
)

func NewRouter(opts Options) {
	clients.New(opts.Router, opts.Ctrl)

	if os.Getenv("ENV") != "production" {
		swagger.New(opts.Router)
	}
}
