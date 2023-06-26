package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	"github.com/windbnb/review-service/model"
	"github.com/windbnb/review-service/service"
	"github.com/windbnb/review-service/tracer"
)

type Handler struct {
	Service *service.RatingService
	Tracer  opentracing.Tracer
	Closer  io.Closer
}

func (handler *Handler) RateHost(w http.ResponseWriter, r *http.Request) {
	span := tracer.StartSpanFromRequest("hostRatingHandler", handler.Tracer, r)
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling rating host at %s\n", r.URL.Path)),
	)

	var hostRatingRequest model.HostRatingRequest
	json.NewDecoder(r.Body).Decode(&hostRatingRequest)

	ctx := tracer.ContextWithSpan(context.Background(), span)
	hostRating, err := handler.Service.SaveHostRating(&hostRatingRequest, ctx)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		tracer.LogError(span, err)
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusUnauthorized})
		return
	}

	json.NewEncoder(w).Encode(hostRating)
}

func (handler *Handler) GetHostAverageReview(w http.ResponseWriter, r *http.Request) {
	span := tracer.StartSpanFromRequest("hostRatingHandler", handler.Tracer, r)
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling rating host at %s\n", r.URL.Path)),
	)

	params := mux.Vars(r)
	hostId, _ := strconv.ParseUint(params["id"], 10, 32)

	ctx := tracer.ContextWithSpan(context.Background(), span)
	hostRating, err := handler.Service.GetAverageHostRating(uint(hostId), ctx)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		tracer.LogError(span, err)
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusUnauthorized})
		return
	}

	json.NewEncoder(w).Encode(hostRating)
}

func (handler *Handler) RateAccomodation(w http.ResponseWriter, r *http.Request) {
	span := tracer.StartSpanFromRequest("accomodationRatingHandler", handler.Tracer, r)
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling rating accomodation at %s\n", r.URL.Path)),
	)

	var accomodationRatingRequest model.AccomodationRatingRequest
	json.NewDecoder(r.Body).Decode(&accomodationRatingRequest)

	ctx := tracer.ContextWithSpan(context.Background(), span)
	accomodationRating, err := handler.Service.SaveAccomodationRating(&accomodationRatingRequest, ctx)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		tracer.LogError(span, err)
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusUnauthorized})
		return
	}

	json.NewEncoder(w).Encode(accomodationRating)
}

func (handler *Handler) GetAccomodationAverageReview(w http.ResponseWriter, r *http.Request) {
	span := tracer.StartSpanFromRequest("accomodationAverageReviewHandler", handler.Tracer, r)
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling average accomodation at %s\n", r.URL.Path)),
	)

	params := mux.Vars(r)
	accomodationId, _ := strconv.ParseUint(params["id"], 10, 32)

	ctx := tracer.ContextWithSpan(context.Background(), span)
	accomodationRating, err := handler.Service.GetAverageAccomodationRating(uint(accomodationId), ctx)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		tracer.LogError(span, err)
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusUnauthorized})
		return
	}

	json.NewEncoder(w).Encode(accomodationRating)
}

