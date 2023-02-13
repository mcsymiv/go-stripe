package handlers

import (
	"net/http"

	"github.com/mcsymiv/go-stripe/internal/web/config"
	"github.com/mcsymiv/go-stripe/internal/web/render"
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

func New(r *Repository) {
	Repo = r
}

func (repo *Repository) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	repo.App.InfoLog.Println("hit virtual terminal")
	if err := render.Template(w, r, "terminal", nil); err != nil {
		repo.App.ErrorLog.Println("unable to render terminal page template", err)
	}
}

func (repo *Repository) PaymentSucceeded(w http.ResponseWriter, r *http.Request) {
	repo.App.InfoLog.Println("hit payment succeeded")

	err := r.ParseForm()
	if err != nil {
		repo.App.ErrorLog.Println("unable to parse payment form", err)

		return
	}

	email := r.Form.Get("email")
	hn := r.Form.Get("holder_name")
	pi := r.Form.Get("payment_intent")
	pm := r.Form.Get("payment_method")
	pa := r.Form.Get("payment_amount")
	pc := r.Form.Get("payment_currency")

	data := make(map[string]interface{})
	data["hn"] = hn
	data["pi"] = pi
	data["pm"] = pm
	data["pa"] = pa
	data["pc"] = pc
	data["email"] = email

	err = render.Template(w, r, "succeeded", &render.TemplateData{
		Data: data,
	})
	if err != nil {
		repo.App.ErrorLog.Println("unable to render succeeded tmpl", err)
		http.Redirect(w, r, "/virtual-terminal", http.StatusSeeOther)

		return
	}
}
