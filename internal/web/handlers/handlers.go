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
