package controllers

import (
	"github.com/patrickchagastavares/rinha-backend-2024/internal/controllers/clients"
	"github.com/patrickchagastavares/rinha-backend-2024/internal/services"
	"github.com/patrickchagastavares/rinha-backend-2024/pkg/logger"
)

type (
	Container struct {
		Clients clients.IController
	}

	Options struct {
		Srv *services.Container
		Log logger.Logger
	}
)

func New(opts Options) *Container {
	return &Container{
		Clients: clients.New(opts.Srv, opts.Log),
	}
}
