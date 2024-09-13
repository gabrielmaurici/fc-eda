package gateway

import "github.com.br/devfullcycle/fc-ms-wallet/internal/entity"

type AccountGateway interface {
	FindByID(id string) (*entity.Account, error)
	UpdateBalance(account *entity.Account) error
}
