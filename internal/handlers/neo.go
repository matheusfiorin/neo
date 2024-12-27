package handlers

import (
	"encoding/json"
	"go-rest-api/internal/services"
	"net/http"
	"time"
)

type NEOHandler struct {
	nasaService *services.NASAService
}

func NewNEOHandler(nasaService *services.NASAService) *NEOHandler {
	return &NEOHandler{nasaService: nasaService}
}

func (h *NEOHandler) GetNEOs(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	if startDate == "" {
		startDate = time.Now().Format("2006-01-02")
	}
	if endDate == "" {
		endDate = time.Now().AddDate(0, 0, 7).Format("2006-01-02")
	}

	neos, err := h.nasaService.GetNEOs(startDate, endDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(neos)
}
