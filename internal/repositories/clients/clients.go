package clients

import (
	"context"
	"errors"

	"github.com/patrickchagastavares/rinha-backend-2024/internal/entities"
)

type IRepository interface {
	CreateTransaction(ctx context.Context, transaction *entities.TransactionRequest, limit int) (int, error)
	FindExtract(ctx context.Context, id int) ([]entities.Transaction, error)

	IsErrNotLimit(err error) bool
	// IsErrDuplicate(err error) bool
}

var (
	// transaction_operation(transaction_id uuid, cliente_id integer, value integer,
	// tipo transactiontype, description character varying,
	// transaction_created_at timestamp with time zone, limite integer)
	createTransaction = `select transaction_operation($1,$2,$3,$4,$5,$6,$7);`
	findTransactions  = `
	select t.value, t."type", t.description, t.last_balance, t.created_at 
	from transactions t
	where t.client_id = $1
	order by t.created_at desc
	limit 10;`

	// Errors
	ErrTransaction = errors.New("problem to create transactin")
	ErrLimit       = errors.New("no limit available")
)
