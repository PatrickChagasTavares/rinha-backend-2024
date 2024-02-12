package services

import (
	"github.com/patrickchagastavares/rinha-backend-2024/internal/repositories"
	"github.com/patrickchagastavares/rinha-backend-2024/internal/services/clients"
	"github.com/patrickchagastavares/rinha-backend-2024/pkg/logger"
)

type (
	Container struct {
		Clients clients.IService
	}

	Options struct {
		Repo *repositories.Container
		Log  logger.Logger
	}
)

func New(opts Options) *Container {
	return &Container{
		Clients: clients.New(opts.Repo, opts.Log),
	}
}
