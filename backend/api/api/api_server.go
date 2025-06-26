package api

import (
	"api/api/handler"
	"fmt"
	"net/http"
)

type APIServer struct {
	server *http.Server
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight (OPTIONS) requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}


func setupServerMux() *http.ServeMux {
  apiV1 := http.NewServeMux()

  uniAPI := handler.NewUniversityHandler().GetRoutes()
  uniMux := http.NewServeMux()
  uniMux.Handle("/universities/", http.StripPrefix("/universities", uniAPI))

  apiV1.Handle("/api/v1/", http.StripPrefix("/api/v1", uniMux))

  return apiV1
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		server: &http.Server{
			Addr:    addr,
			Handler: corsMiddleware(setupServerMux()),
		},
	}

}

func (api *APIServer) Run() error {
	if err := api.server.ListenAndServe(); err != nil {
		return fmt.Errorf("Unable to run server: %s", err)
	}

	return nil
}
