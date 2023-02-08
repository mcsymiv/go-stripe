package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mcsymiv/go-stripe/internal/api/config"
)

var Repo *Repository

type Repository struct {
	App *config.Application
}

func NewRepository(a *config.Application) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// FE client stripe payload sent
type stripePayload struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

// BE response to FE client
type responce struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
	Content string `json:"content"`
	Id      string `json:"id"`
}

func (repo *Repository) GetPaymentIntent(w http.ResponseWriter, r *http.Request) {
	jres := responce{
		Ok: true,
	}

	jout, err := json.MarshalIndent(jres, "", "\t")
	if err != nil {
		repo.App.ErrorLog.Println("unable to marshall json", err)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jout)
}
