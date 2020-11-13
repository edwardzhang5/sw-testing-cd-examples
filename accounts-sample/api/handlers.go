package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/drbyronw/accounts/models"
	"github.com/go-chi/chi"
)

// Home handler.
func (wa *WebApp) Home(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("BANK Accounts API v: %s", os.Getenv("API_VERSION"))
	if _, err := w.Write([]byte(msg)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// GetAccount retreives an account given an accountID
func (wa *WebApp) GetAccount(w http.ResponseWriter, r *http.Request) {
	var err error

	accountID := chi.URLParam(r, "accountID")
	acct, err := wa.Accounts.Find(accountID)
	if err != nil {
		msg := fmt.Sprintf("Unable to process request: %v", err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	respondWithJSON(w, http.StatusOK, acct)
}

// DepositFunds enables cash to be deposited to the account
func (wa *WebApp) DepositFunds(w http.ResponseWriter, r *http.Request) {
	var err error
	data := make(map[string]interface{})
	account := chi.URLParam(r, "accountID")

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		msg := fmt.Sprintf("[DepositFunds]: unable to parse request - %v", err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	amt := data["amount"].(float64)

	newBalance, err := wa.Accounts.Deposit(account, amt)
	if err != nil {
		msg := fmt.Sprintf("unable to process deposit:   %v", err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	respondWithJSON(w, http.StatusAccepted, map[string]interface{}{
		"message":     "deposit successful",
		"new_balance": newBalance.String(),
	})

}

// TransferCash provides the handler to transfer funds between accounts
func (wa *WebApp) TransferCash(w http.ResponseWriter, r *http.Request) {
	var err error
	data := make(map[string]interface{})

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		msg := fmt.Sprintf("unable to parse request:   %v", err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	from := data["from"].(string)
	to := data["to"].(string)
	amount := data["amount"].(float64)
	err = wa.Accounts.Transfer(from, to, amount)
	if err != nil {
		msg := fmt.Sprintf("unable to process transfer:   %v", err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	respondWithJSON(w, http.StatusAccepted, map[string]interface{}{
		"message": "transfer completed successfully",
	})
}

// GetClientAccounts accepts a 'client_id' as a query parametert to find all
// accounts for that client ID
func (wa *WebApp) GetClientAccounts(w http.ResponseWriter, r *http.Request) {
	var accts []models.Account
	var err error
	clientID := chi.URLParam(r, "clientID")
	accts, err = wa.Accounts.FindAll(clientID)
	if err != nil {
		msg := fmt.Sprintf("Unable to process request: %v", err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	respondWithJSON(w, http.StatusOK, accts)
}

// PostClientAccounts accepts a 'client_id' json key to find all
// accounts for that client ID
func (wa *WebApp) PostClientAccounts(w http.ResponseWriter, r *http.Request) {
	var accts []models.Account
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		msg := fmt.Sprintf("Unable to parse request: %v", err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	clientID := data["client_id"].(string)
	accts, err = wa.Accounts.FindAll(clientID)
	if err != nil {
		msg := fmt.Sprintf("Unable to process request: %v", err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	respondWithJSON(w, http.StatusOK, accts)
}

// AddAccount handler accepts a JSON Post of account info
func (wa *WebApp) AddAccount(w http.ResponseWriter, r *http.Request) {
	var acct models.Account
	err := json.NewDecoder(r.Body).Decode(&acct)
	if err != nil {
		msg := fmt.Sprintf("Unable to parse request: %v", err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	if acct.ClientID == "" || acct.Type == "" {
		msg := "Unable to parse request - incomplete Account info"
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	err = wa.Accounts.Create(acct)
	if err != nil {
		msg := fmt.Sprintf("Unable to process create account: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "account successfully created",
	})
}

// Ping responds to request to confirm access
func (wa *WebApp) Ping(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

// respondWithJSON write json response format
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if _, err := w.Write(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
