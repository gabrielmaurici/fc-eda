package update_account_balance

import (
	"github.com.br/devfullcycle/fc-ms-wallet/internal/gateway"
)

type UpdateAccountBalanceInputDto struct {
	AccountId string  `json:"account_id"`
	Amount    float64 `json:"amount"`
}

type UpdateAccountBalanceOutputDto struct {
	AccountId string  `json:"account_id"`
	Amount    float64 `json:"amount"`
}

type UpdateAccountBalanceUseCase struct {
	AccountGateway gateway.AccountGateway
}

func NewUpdateAccountBalanceUseCase(a gateway.AccountGateway) *UpdateAccountBalanceUseCase {
	return &UpdateAccountBalanceUseCase{
		AccountGateway: a,
	}
}

func (uc *UpdateAccountBalanceUseCase) Execute(input UpdateAccountBalanceInputDto) error {
	account, err := uc.AccountGateway.FindByID(input.AccountId)
	if err != nil {
		return err
	}
	account.UpdateBalance(input.Amount)
	err = uc.AccountGateway.UpdateBalance(account)
	if err != nil {
		return err
	}
	return nil
}
