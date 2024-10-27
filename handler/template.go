package handler

import "net/http"

type templateVariables map[string]any

func (h *Handler) renderTemplate(
	w http.ResponseWriter, name string, variables templateVariables,
) error {
	return h.templates.ExecuteTemplate(w, name, variables)
}
