package handlers

import (
	"net/http"
	"time"

	"github.com/mcsymiv/go-stripe/internal/models"
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
	stringMap := make(map[string]string)
	stringMap["pb_test"] = repo.App.Config.Stripe.Key

	err := render.Template(w, r, "terminal", &render.TemplateData{
		StringMap: stringMap,
	}, "form-js")

	if err != nil {
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

func (repo *Repository) ChargeItem(w http.ResponseWriter, r *http.Request) {
	repo.App.InfoLog.Println("hit charge item page")

	wicker := models.Wicker{
		Id:          1,
		Name:        "fine basket wicker",
		ImageName:   "wicker",
		Price:       1200,
		Description: "handmade fine wicker basket",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	data := make(map[string]interface{})
	data["wicker"] = wicker

	err := render.Template(w, r, "wicker", &render.TemplateData{
		Data: data,
	}, "form-js")

	if err != nil {
		repo.App.ErrorLog.Printf("unable to render wicker page. Error: %v", err)
	}
}
