package general

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HelloResponse struct {
	Message string `json:"message"`
}

// helloHandler handles the /hello route
func helloHandler(w http.ResponseWriter, r *http.Request) {
	resp := HelloResponse{Message: "Hello, Mehul! 👋"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func serverExample() {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Register handler
	r.Get("/hello", helloHandler)

	http.ListenAndServe(":8080", r)
}
