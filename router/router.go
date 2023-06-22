package router

import (
	"github.com/gorilla/mux"
	"github.com/windbnb/review-service/handler"
	"github.com/windbnb/review-service/metrics"
)

func ConfigureRouter(handler *handler.Handler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/host-rating", handler.RateHost).Methods("POST")
	router.HandleFunc("/api/host-rating/host/{id}", handler.TestHandler).Methods("GET")
	router.HandleFunc("/api/accomodation-rating", handler.TestHandler).Methods("POST")
	router.HandleFunc("/api/accomodation-rating/accomodation/{id}", handler.TestHandler).Methods("GET")

	router.Path("/metrics").Handler(metrics.MetricsHandler())

	return router
}
