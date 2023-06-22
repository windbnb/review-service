package service

import (
	"context"
	"os"

	"github.com/windbnb/review-service/model"
	"github.com/windbnb/review-service/repository"
	"github.com/windbnb/review-service/tracer"
)

var ratingServiceUrl = os.Getenv("ratingServiceUrl") + "/api/accomodation/"

type RatingService struct {
	Repo repository.IRepository
}

func (s *RatingService) SaveHostRating(hostRatingRequest *model.HostRatingRequest, ctx context.Context) (*model.HostRating, error) {
	span := tracer.StartSpanFromContext(ctx, "saveHostRatingService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	// Checks

	// Get if had accomodation

	var hostRatingRequestData = model.HostRating{
		GuestId: hostRatingRequest.GuestId,
		HostId:  hostRatingRequest.HostId,
		Raiting: hostRatingRequest.Raiting}

	s.Repo.RateHost(&hostRatingRequestData, ctx)

	return &hostRatingRequestData, nil
}

func (s *RatingService) DummyService(ctx context.Context) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "saveHostRatingService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	return "test radi", nil
}
