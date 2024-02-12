package entities

import (
	"time"

	"github.com/google/uuid"
)

const (
	TipoDebito  Tipo = "c"
	TipoCredito Tipo = "d"
)

type (
	Tipo string

	Transaction struct {
		Valor       uint64    `db:"value" json:"valor"`
		Tipo        Tipo      `db:"type" json:"tipo"`
		Description string    `db:"description" json:"descricao"`
		Balance     int       `db:"last_balance" json:"-"`
		CreatedAt   time.Time `db:"created_at" json:"realizada_em"`
	}

	TransactionRequest struct {
		ID          string    `json:"-"`
		ClientID    int       `json:"-"`
		Value       int       `json:"valor" validate:"required,min=0"`
		Type        Tipo      `json:"tipo" validate:"required,typeTransaction"`
		Description string    `json:"descricao" validate:"required,max=10"`
		CreatedAt   time.Time `json:"-"`
	}

	TransactionBalance struct {
		Limit   int `json:"limite"`
		Balance int `json:"saldo"`
	}
	Balance struct {
		Total    int       `json:"total"`
		Limit    int       `json:"limite"`
		CreateAt time.Time `json:"data_extrato"`
	}

	Extract struct {
		Balance         Balance       `json:"saldo"`
		LastTransaction []Transaction `json:"ultimas_transacoes"`
	}
)

func (tr *TransactionRequest) PreSave() {
	tr.ID = uuid.NewString()
	tr.CreatedAt = time.Now().UTC()
}
