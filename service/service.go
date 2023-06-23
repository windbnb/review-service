package service

import (
	"context"
	"errors"
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

	_, err := s.Repo.DeleteGuestHostRatings(hostRatingRequest.GuestId, hostRatingRequest.HostId, ctx)
	if err != nil {
		return nil, err
	}

	// Checks
	if hostRatingRequest.Raiting < 1 || hostRatingRequest.Raiting > 5 {
		return nil, errors.New("rating must be between 1 and 5")
	}

	// Get if had accomodation

	var hostRatingRequestData = model.HostRating{
		GuestId: hostRatingRequest.GuestId,
		HostId:  hostRatingRequest.HostId,
		Raiting: hostRatingRequest.Raiting}

	s.Repo.RateHost(&hostRatingRequestData, ctx)

	return &hostRatingRequestData, nil
}

func (s *RatingService) SaveAccomodationRating(accomodationRatingRequest *model.AccomodationRatingRequest, ctx context.Context) (*model.AccomodationRating, error) {
	span := tracer.StartSpanFromContext(ctx, "saveAccomodationRatingService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	_, err := s.Repo.DeleteGuestAccomodationRatings(accomodationRatingRequest.GuestId, accomodationRatingRequest.AccomodationId, ctx)
	if err != nil {
		return nil, err
	}

	// Checks
	if accomodationRatingRequest.Raiting < 1 || accomodationRatingRequest.Raiting > 5 {
		return nil, errors.New("rating must be between 1 and 5")
	}

	// Get if had accomodation

	var accomodationRatingRequestData = model.AccomodationRating{
		GuestId:        accomodationRatingRequest.GuestId,
		AccomodationId: accomodationRatingRequest.AccomodationId,
		Raiting:        accomodationRatingRequest.Raiting}

	s.Repo.RateAccomodation(&accomodationRatingRequestData, ctx)

	return &accomodationRatingRequestData, nil
}

func (s *RatingService) GetAverageHostRating(hostId uint, ctx context.Context) (*model.HostAvgRating, error) {
	span := tracer.StartSpanFromContext(ctx, "getAverageHostRatingService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	ratings, err := s.Repo.FindAllHostRatings(uint(hostId), ctx)
	if err != nil {
		return nil, err
	}

	var avgRating float32 = 0
	for _, rating := range *ratings {
		avgRating += float32(rating.Raiting)
	}

	if len(*ratings) > 0 {
		avgRating /= float32(len(*ratings))
	}

	var result = model.HostAvgRating{
		HostId:         uint(hostId),
		AverageRaiting: avgRating,
	}

	return &result, nil
}

func (s *RatingService) GetAverageAccomodationRating(accomodationId uint, ctx context.Context) (*model.AccomodationAvgRating, error) {
	span := tracer.StartSpanFromContext(ctx, "getAverageAccomodationRatingService")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	ratings, err := s.Repo.FindAllAccomodationRatings(uint(accomodationId), ctx)
	if err != nil {
		return nil, err
	}

	var avgRating float32 = 0
	for _, rating := range *ratings {
		avgRating += float32(rating.Raiting)

	}

	if len(*ratings) > 0 {
		avgRating /= float32(len(*ratings))
	}

	var result = model.AccomodationAvgRating{
		AccomodationId: uint(accomodationId),
		AverageRaiting: avgRating,
	}

	return &result, nil
}

func (s *RatingService) DummyService(ctx context.Context) (string, error) {
	span := tracer.StartSpanFromContext(ctx, "saveHostRatingService")
	defer span.Finish()

	tracer.ContextWithSpan(context.Background(), span)

	return "test radi", nil
}
