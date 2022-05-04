package handlers

import (
	"net/http"

	"github.com/Marvellous-Chimaraoke/bookings/pkg/config"
	"github.com/Marvellous-Chimaraoke/bookings/pkg/models"
	"github.com/Marvellous-Chimaraoke/bookings/pkg/render"
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// Repo is the repository used by the handlers
var Repo *Repository

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	res := map[string]string{
		"data": "Marvellous'",
	}
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{StringMap: res})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello world V2"

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
