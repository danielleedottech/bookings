package handlers

import (
	"fmt"
	"github.com/danielleedottech/bookings/config"
	"github.com/danielleedottech/bookings/models"
	"github.com/danielleedottech/bookings/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "home.gohtml", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "HelloMap"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	fmt.Println(remoteIP)
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}
