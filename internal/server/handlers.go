package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func mapHandlers() *mux.Router {
	// setup router
	r := mux.NewRouter()

	v1Router := r.PathPrefix("/api/v1").Subrouter()
	v1Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
		if err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
			return
		}
	}).Methods("GET")

	return r
}
