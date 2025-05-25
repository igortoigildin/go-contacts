package subscriberrepository

import (
	"github.com/igortoigildin/go-contacts/subscriber/pkg/postgres/transaction_manager"
)

type Repository struct {
	txManager transaction_manager.TransactionManagerAPI
}

func NewRepository(txManager transaction_manager.TransactionManagerAPI) *Repository {
	return &Repository{
		txManager: txManager,
	}
}
