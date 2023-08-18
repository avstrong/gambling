package web

import (
	"encoding/json"
	"net/http"

	"emperror.dev/errors"
	"github.com/avstrong/gambling/internal/uservice"
	"github.com/gorilla/mux"
)

func (s *Server) depositHandler(w http.ResponseWriter, r *http.Request) {
	var input uservice.DepositInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	if err = input.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	_, err = s.uService.Deposit(r.Context(), &input)
	if errors.Is(err, uservice.ErrNotFound) {
		http.NotFound(w, r)

		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}

func (s *Server) withdrawHandler(w http.ResponseWriter, r *http.Request) {
	var input uservice.WithdrawInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	if err = input.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	_, err = s.uService.Withdraw(r.Context(), &input)
	if errors.Is(err, uservice.ErrNotFound) {
		http.NotFound(w, r)

		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}

func (s *Server) balanceHandler(w http.ResponseWriter, r *http.Request) {
	var input uservice.BalanceInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	if err = input.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	_, err = s.uService.Balance(r.Context(), &input)
	if errors.Is(err, uservice.ErrNotFound) {
		http.NotFound(w, r)

		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}

func (s *Server) addRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc("/api/v1/wallet/deposit", s.depositHandler)
	r.HandleFunc("/api/v1/wallet/withdraw", s.withdrawHandler)
	r.HandleFunc("/api/v1/wallet/balance/:user_id", s.balanceHandler)

	r.Use(recoverMiddleware())
	r.Use(loggerMiddleware())

	return r
}
