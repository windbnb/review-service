package router

import (
	"github.com/gorilla/mux"
	"github.com/windbnb/review-service/handler"
	"github.com/windbnb/review-service/metrics"
)

func ConfigureRouter(handler *handler.Handler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/host-rating", metrics.MetricProxy(handler.RateHost)).Methods("POST")
	router.HandleFunc("/api/host-rating/host/{id}", metrics.MetricProxy(handler.GetHostAverageReview)).Methods("GET")
	router.HandleFunc("/api/accomodation-rating", metrics.MetricProxy(handler.RateAccomodation)).Methods("POST")
	router.HandleFunc("/api/accomodation-rating/accomodation/{id}", metrics.MetricProxy(handler.GetAccomodationAverageReview)).Methods("GET")

	router.Path("/metrics").Handler(metrics.MetricsHandler())

	return router
}
