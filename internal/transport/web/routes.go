//nolint:dupl
package web

import (
	"encoding/json"
	"net/http"

	"emperror.dev/errors"
	"github.com/avstrong/gambling/internal/gmanager"
	"github.com/avstrong/gambling/internal/uservice"
	"github.com/gorilla/mux"
)

func (s *Server) createGameHandler(w http.ResponseWriter, r *http.Request) {
	var input gmanager.CreateGameInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	if err = input.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	out, err := s.gManager.CreateGame(r.Context(), &input)
	if err != nil {
		s.lg.Err(err).Msg("Cannot create a game")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(out); err != nil {
		s.lg.Err(err).Msg("Cannot encode result of creating game")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}
}

func (s *Server) registerPlayersHandler(w http.ResponseWriter, r *http.Request) {
	var input gmanager.RegisterPlayerInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	if err = input.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	out, err := s.gManager.RegisterPlayer(r.Context(), &input)
	if errors.Is(err, gmanager.ErrInsufficientFunds) {
		http.Error(w, "deposit money", http.StatusForbidden)

		return
	}

	if errors.Is(err, gmanager.ErrAlreadyExists) {
		w.WriteHeader(http.StatusNoContent)

		return
	}

	if err != nil {
		s.lg.Err(err).Msg("Cannot register a player")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(out); err != nil {
		s.lg.Err(err).Msg("Cannot encode result of player registration")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}
}

func (s *Server) playGameHandler(w http.ResponseWriter, r *http.Request) {
	var input gmanager.StartGameInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	if err = input.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	out, err := s.gManager.Play(r.Context(), &input)
	if err != nil {
		s.lg.Err(err).Msg("Cannot play a game")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(out); err != nil {
		s.lg.Err(err).Msg("Cannot encode result of launching game")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}
}

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
		s.lg.Err(err).Msg("Cannot deposit money")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

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
		s.lg.Err(err).Msg("Cannot withdraw money")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

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

	out, err := s.uService.Balance(r.Context(), &input)
	if errors.Is(err, uservice.ErrNotFound) {
		http.NotFound(w, r)

		return
	}

	if err != nil {
		s.lg.Err(err).Msg("Cannot get balance")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(out); err != nil {
		s.lg.Err(err).Msg("Cannot encode result of getting balance")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}
}

func (s *Server) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var input uservice.RegisterUserInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	if err = input.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	out, err := s.uService.RegisterUser(r.Context(), &input)
	if errors.Is(err, uservice.ErrAlreadyExists) {
		w.WriteHeader(http.StatusNoContent)

		return
	}

	if err != nil {
		s.lg.Err(err).Msg("Cannot register user")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(out); err != nil {
		s.lg.Err(err).Msg("Cannot encode result of register user")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}
}

func (s *Server) addRoutes(r *mux.Router) *mux.Router {
	r.HandleFunc("/api/v1/game", s.createGameHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/game/register", s.registerPlayersHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/game/play", s.playGameHandler).Methods(http.MethodPost)

	r.HandleFunc("/api/v1/user", s.registerUserHandler).Methods(http.MethodPost)

	r.HandleFunc("/api/v1/wallet/deposit", s.depositHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/wallet/withdraw", s.withdrawHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/wallet/balance", s.balanceHandler).Methods(http.MethodGet)

	r.Use(recoverMiddleware())
	r.Use(loggerMiddleware())

	return r
}
