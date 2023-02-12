package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/mcsymiv/go-stripe/internal/api/config"
	"github.com/mcsymiv/go-stripe/internal/card"
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
	Message string `json:"message,omitempty"`
	Content string `json:"content,omitempty"`
	Id      string `json:"id,omitempty"`
}

func (repo *Repository) PaymentIntent(w http.ResponseWriter, r *http.Request) {
	var payload stripePayload

	/// read and parse request body
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		repo.App.ErrorLog.Println("unable to parse request body", err)

		return
	}

	amount, err := strconv.Atoi(payload.Amount)
	if err != nil {
		repo.App.ErrorLog.Println("unable to convert payload amount", err)

		return
	}

	c := card.Card{
		Secret:   repo.App.Config.Stripe.Secret,
		Key:      repo.App.Config.Stripe.Key,
		Currency: payload.Currency,
	}

	pi, msg, err := c.Charge(payload.Currency, amount)
	if err != nil {
		repo.App.ErrorLog.Println("unable to charge card", msg, err)

		eres := responce{
			Ok:      false,
			Message: msg,
			Content: "from card charge error",
		}

		erout, err := json.MarshalIndent(eres, "", "\t")
		if err != nil {
			repo.App.ErrorLog.Println("unable to marshall error response", err)

			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(erout)

		return
	}

	jout, err := json.MarshalIndent(pi, "", "\t")
	if err != nil {
		repo.App.ErrorLog.Println("unable to marshall payment intent", err)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jout)
}
