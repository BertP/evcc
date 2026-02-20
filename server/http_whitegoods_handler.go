package server

import (
	"net/http"

	"github.com/evcc-io/evcc/core/whitegood"
	"github.com/gorilla/mux"
)

// whitegoodsHandler returns the current appliances and their state
func whitegoodsHandler(wgCoord *whitegood.Coordinator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := make(map[string]any)
		res["appliances"] = wgCoord.Appliances()
		jsonWrite(w, res)
	}
}

// whitegoodStartHandler handles manual forced start
func whitegoodStartHandler(wgCoord *whitegood.Coordinator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]

		err := wgCoord.StartAppliance(name)
		if err != nil {
			jsonError(w, http.StatusBadRequest, err)
			return
		}

		res := struct {
			Started bool `json:"started"`
		}{
			Started: true,
		}

		jsonWrite(w, res)
	}
}

// RegisterWhitegoodsHandler binds the whitegoods endpoints
func (s *HTTPd) RegisterWhitegoodsHandler(wgCoord *whitegood.Coordinator) {
	router := s.Server.Handler.(*mux.Router)
	api := router.PathPrefix("/api").Subrouter()

	api.HandleFunc("/whitegoods", whitegoodsHandler(wgCoord)).Methods(http.MethodGet)
	api.HandleFunc("/whitegoods/{name:[a-zA-Z0-9_.-]+}/start", whitegoodStartHandler(wgCoord)).Methods(http.MethodPost)
}
