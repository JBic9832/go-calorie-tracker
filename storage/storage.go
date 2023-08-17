package storage

import "github.com/jbic9832/go-calorie-tracker/types"

type Storage interface {
	// Account Operations
	CreateAccount(*types.Account) error
    DeleteAccount(int) error
    UpdateAccount(*types.Account) error
    GetAccount(int) (*types.Account, error)
}
