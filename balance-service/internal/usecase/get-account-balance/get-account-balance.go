package get_account_balance

import (
	"github.com.br/devfullcycle/fc-ms-wallet/internal/gateway"
)

type GetAccountBalanceInputDto struct {
	AccountId string `json:"account_id"`
}

type GetAccountBalanceOutputDto struct {
	Balance float64 `json:"balance"`
}

type GetAccountBalanceUseCase struct {
	AccountGateway gateway.AccountGateway
}

func NewGetAccountBalanceUseCase(a gateway.AccountGateway) *GetAccountBalanceUseCase {
	return &GetAccountBalanceUseCase{
		AccountGateway: a,
	}
}

func (uc *GetAccountBalanceUseCase) Execute(input GetAccountBalanceInputDto) (output *GetAccountBalanceOutputDto, err error) {
	account, err := uc.AccountGateway.FindByID(input.AccountId)
	if err != nil {
		return nil, err
	}
	return &GetAccountBalanceOutputDto{
		Balance: account.Balance,
	}, nil
}
