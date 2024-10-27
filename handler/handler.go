package handler

import (
	"html/template"
	"log"
	"net/http"
	"steamserverlauncher/handler/middleware"
	"steamserverlauncher/sourcequery"
)

type Handler struct {
	templates     *template.Template
	client        *sourcequery.SourceQueryIntegration
	serverAddress string
	serverName    string
}

func New(
	client *sourcequery.SourceQueryIntegration,
	templatesPattern, serverAddress, serverName string,
) (*Handler, error) {
	templates, err := template.ParseGlob(templatesPattern)
	if err != nil {
		return nil, err
	}

	return &Handler{
		templates:     templates,
		client:        client,
		serverAddress: serverAddress,
		serverName:    serverName,
	}, nil
}

func (h *Handler) RegisterRoutes() http.Handler {
	mux := http.DefaultServeMux

	staticFileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", staticFileServer))

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		_ = h.renderTemplate(w, "base", templateVariables{
			"ServerAddress": h.serverAddress,
			"ServerName":    h.serverName,
		})
	})

	mux.HandleFunc("GET /players", func(w http.ResponseWriter, r *http.Request) {
		currentPlayers, err := h.client.CurrentPlayers(r.Context())
		if err != nil {
			log.Println("h.client.CurrentPlayers:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_ = h.renderTemplate(w, "players", templateVariables{
			"Players": currentPlayers,
		})
	})

	wrappedMux := middleware.NewLogger(mux)
	return wrappedMux
}
