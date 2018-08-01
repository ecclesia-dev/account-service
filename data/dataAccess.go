package data

import (
	"github.com/ecclesia-dev/account-service/models"
)

type DataAccess interface {
	CreateAccount(models.Account) error
	GetAllAccounts() ([]models.Account, error)
	GetAccountByID(string) (models.Account, error)
	GetAccountByEmail(string) (models.Account, error)
	UpdateAccount(string, map[string]interface{}) error
	RemoveAccount(string) error
}
