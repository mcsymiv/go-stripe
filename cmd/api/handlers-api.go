package main

import (
	"encoding/json"
	"net/http"
)

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

func (a *app) GetPaymentIntent(w http.ResponseWriter, r *http.Request) {
	jres := responce{
		Ok: true,
	}

	jout, err := json.MarshalIndent(jres, "", "\t")
	if err != nil {
		a.errorLog.Println("unable to marshall json", err)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jout)
}
