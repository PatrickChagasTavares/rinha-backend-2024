package clients

import (
	"context"
	"time"

	"github.com/patrickchagastavares/rinha-backend-2024/internal/entities"
	"github.com/patrickchagastavares/rinha-backend-2024/internal/repositories"
	"github.com/patrickchagastavares/rinha-backend-2024/pkg/logger"
)

type (
	IService interface {
		CreateTransaction(ctx context.Context, transaction *entities.TransactionRequest) (*entities.TransactionBalance, error)
		FindExtract(ctx context.Context, id int) (extract entities.Extract, err error)
	}

	services struct {
		repositories *repositories.Container
		log          logger.Logger
	}
)

func New(repo *repositories.Container, log logger.Logger) IService {
	return &services{repositories: repo, log: log}
}

func (srv *services) CreateTransaction(ctx context.Context, transaction *entities.TransactionRequest) (*entities.TransactionBalance, error) {
	limit, ok := clientsAvailable[transaction.ClientID]
	if !ok {
		return nil, ErrClientNotFound
	}

	transaction.PreSave()

	balance, err := srv.repositories.Database.Clients.CreateTransaction(ctx, transaction, limit)
	if err != nil {
		if srv.repositories.Database.Clients.IsErrNotLimit(err) {
			return nil, ErrClientNotLimit
		}
		return nil, err
	}

	return &entities.TransactionBalance{
		Limit:   limit,
		Balance: balance,
	}, nil
}

func (srv *services) FindExtract(ctx context.Context, id int) (extract entities.Extract, err error) {
	limit, ok := clientsAvailable[id]
	if !ok {
		err = ErrClientNotFound
		return
	}

	transactions, err := srv.repositories.Database.Clients.FindExtract(ctx, id)
	if err != nil {
		return
	}

	extract.Balance = entities.Balance{
		Limit:    limit,
		CreateAt: time.Now().UTC(),
	}
	extract.LastTransaction = transactions

	if len(transactions) > 0 {
		extract.Balance.Total = transactions[0].Balance
	}

	return
}
