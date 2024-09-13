package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	get_account_balance "github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/get-account-balance"
	"github.com/go-chi/chi/v5"
)

type WebAccountHandler struct {
	GetAccountBalanceUseCase get_account_balance.GetAccountBalanceUseCase
}

func NewWebAccountHandler(getAccountBalanceUseCase get_account_balance.GetAccountBalanceUseCase) *WebAccountHandler {
	return &WebAccountHandler{
		GetAccountBalanceUseCase: getAccountBalanceUseCase,
	}
}

func (h *WebAccountHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "account_id")
	getBalanceInput := get_account_balance.GetAccountBalanceInputDto{
		AccountId: id,
	}
	output, err := h.GetAccountBalanceUseCase.Execute(getBalanceInput)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
