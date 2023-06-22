package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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
		tracer.LogString("handler", fmt.Sprintf("handling login at %s\n", r.URL.Path)),
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

func (handler *Handler) TestHandler(w http.ResponseWriter, r *http.Request) {
	span := tracer.StartSpanFromRequest("hostRatingHandler", handler.Tracer, r)
	defer span.Finish()
	span.LogFields(
		tracer.LogString("handler", fmt.Sprintf("handling login at %s\n", r.URL.Path)),
	)

	var hostRatingRequest model.HostRatingRequest
	json.NewDecoder(r.Body).Decode(&hostRatingRequest)

	// ctx := tracer.ContextWithSpan(context.Background(), span)
	tracer.ContextWithSpan(context.Background(), span)
	// token, err := handler.Service.Login(credentials, ctx)

	w.Header().Set("Content-Type", "application/json")
	// if err != nil {
	// 	tracer.LogError(span, err)
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error(), StatusCode: http.StatusUnauthorized})
	// 	return
	// }

	json.NewEncoder(w).Encode("test")
}
