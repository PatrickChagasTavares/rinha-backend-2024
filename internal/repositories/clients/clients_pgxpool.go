package clients

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/patrickchagastavares/rinha-backend-2024/internal/entities"
	"github.com/patrickchagastavares/rinha-backend-2024/pkg/logger"
)

type repoPgxpool struct {
	log logger.Logger
	db  *pgxpool.Pool
}

func NewPgxPool(log logger.Logger, db *pgxpool.Pool) IRepository {
	return &repoPgxpool{log: log, db: db}
}

func (repo *repoPgxpool) CreateTransaction(ctx context.Context, transaction *entities.TransactionRequest, limit int) (int, error) {
	rows := repo.db.QueryRow(ctx, createTransaction,
		transaction.ID, transaction.ClientID, transaction.Value, transaction.Type,
		transaction.Description, transaction.CreatedAt, limit)

	var balance int
	if err := rows.Scan(&balance); err != nil {
		if validLimit(err) {
			return 0, ErrLimit
		}
		repo.log.ErrorContext(ctx, fmt.Sprintf("clients.repoPgxpool.CreateTransaction query: %v, err: %s", transaction, err.Error()))
		return 0, ErrTransaction
	}
	return balance, nil
}

func (repo *repoPgxpool) FindExtract(ctx context.Context, id int) ([]entities.Transaction, error) {
	rows, err := repo.db.Query(ctx, findTransactions, id)
	if err != nil && err.Error() != sql.ErrNoRows.Error() {
		repo.log.ErrorContext(ctx, fmt.Sprintf("clients.repoPgxpool.FindExtract query: %d, err: %s", id, err.Error()))
		return nil, err
	}

	return scanRows(ctx, rows, repo.scanTransaction)
}

func (repo *repoPgxpool) scanTransaction(ctx context.Context, row pgx.Row) (entities.Transaction, error) {
	transaction := entities.Transaction{}

	err := row.Scan(&transaction.Valor, &transaction.Tipo, &transaction.Description, &transaction.Balance, &transaction.CreatedAt)
	if err != nil {
		return transaction, fmt.Errorf("could not scan cid: %w", err)
	}

	return transaction, nil
}

func (repo *repoPgxpool) IsErrNotLimit(err error) bool {
	return errors.Is(err, ErrLimit)
}

func validLimit(err error) bool {
	return err.Error() == "ERROR: operation is not available: balance is below limit (SQLSTATE P0001)"
}

func scanRows[T entities.Transaction](ctx context.Context, rows pgx.Rows, scanRow func(ctx context.Context, row pgx.Row) (T, error)) ([]T, error) {
	defer rows.Close()
	destRows := make([]T, 0)
	for rows.Next() {
		destRow, err := scanRow(ctx, rows)
		if err != nil {
			return nil, err
		}
		destRows = append(destRows, destRow)
	}
	return destRows, nil
}
