package repositories

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/patrickchagastavares/rinha-backend-2024/internal/repositories/clients"
	"github.com/patrickchagastavares/rinha-backend-2024/pkg/logger"

	_ "github.com/lib/pq"
)

type (
	// Container model to export instance repositories
	Container struct {
		Database SqlContainer
	}

	SqlContainer struct {
		Clients clients.IRepository
	}

	// Options struct of options to create a new repositories
	Options struct {
		DB_URL string
		Log    logger.Logger
	}
)

// New Create a new instance of repositories
func New(opts Options) *Container {
	return &Container{
		Database: SqlContainer{
			Clients: clients.NewPgxPool(opts.Log, newPgxPool(opts.DB_URL)),
		},
	}
}

func newPgxPool(dbURL string) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		panic(fmt.Sprintf("failed to parse connection string: %s", err.Error()))
	}

	db, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		panic(fmt.Sprintf("failed connection with config: %s", err.Error()))
	}

	return db
}
