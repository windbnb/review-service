package service_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/windbnb/review-service/model"
	"github.com/windbnb/review-service/repository"
	"github.com/windbnb/review-service/service"
	"github.com/windbnb/review-service/util"
)

func TestSaveHostRating_Successful_Integration(t *testing.T) {
	db := util.ConnectToMongoDatabase()
	ratingService := &service.RatingService{
		Repo: &repository.Repository{
			Db: db},
	}

	hostRatingRequest := &model.HostRatingRequest{
		GuestId: 1,
		HostId:  1,
		Rating:  4,
	}

	rating, err := ratingService.SaveHostRating(hostRatingRequest, context.Background())

	assert.NotNil(t, rating)
	assert.NoError(t, err)
}

func TestSaveHostRating_Successful(t *testing.T) {
	mockRepo := &MockRepo{
		DeleteGuestHostRatingsFn: func(guestId uint, hostId uint, ctx context.Context) (int64, error) {
			return 1, nil
		},
		RateHostFn: func(hostRating *model.HostRating, ctx context.Context) (*model.HostRating, error) {
			return hostRating, nil
		},
	}

	ratingService := &service.RatingService{
		Repo: mockRepo,
	}

	hostRatingRequest := &model.HostRatingRequest{
		GuestId: 1,
		HostId:  1,
		Rating:  4,
	}

	rating, err := ratingService.SaveHostRating(hostRatingRequest, context.Background())

	assert.NotNil(t, rating)
	assert.NoError(t, err)
}

func TestSaveHostRating_InvalidRating_High(t *testing.T) {
	mockRepo := &MockRepo{
		DeleteGuestHostRatingsFn: func(guestId uint, hostId uint, ctx context.Context) (int64, error) {
			return 1, nil
		},
	}

	ratingService := &service.RatingService{
		Repo: mockRepo,
	}

	hostRatingRequest := &model.HostRatingRequest{
		GuestId: 1,
		HostId:  1,
		Rating:  6, // Invalid rating value
	}

	rating, err := ratingService.SaveHostRating(hostRatingRequest, context.Background())

	assert.Nil(t, rating)
	assert.Error(t, err)
	assert.EqualError(t, err, "rating must be between 1 and 5")
}

func TestSaveHostRating_InvalidRating_Low(t *testing.T) {
	mockRepo := &MockRepo{
		DeleteGuestHostRatingsFn: func(guestId uint, hostId uint, ctx context.Context) (int64, error) {
			return 1, nil
		},
	}

	ratingService := &service.RatingService{
		Repo: mockRepo,
	}

	hostRatingRequest := &model.HostRatingRequest{
		GuestId: 1,
		HostId:  1,
		Rating:  0, // Invalid rating value
	}

	rating, err := ratingService.SaveHostRating(hostRatingRequest, context.Background())

	assert.Nil(t, rating)
	assert.Error(t, err)
	assert.EqualError(t, err, "rating must be between 1 and 5")
}


type MockRepo struct {
	repository.Repository
	DeleteGuestHostRatingsFn func(guestId uint, hostId uint, ctx context.Context) (int64, error)
	RateHostFn func(hostRating *model.HostRating, ctx context.Context) (*model.HostRating, error)
}

func (r *MockRepo) FindGuestHostRatings(guestId uint, hostId uint, ctx context.Context) (*[]model.HostRating, error) {
	return &[]model.HostRating{}, nil

}

func (r *MockRepo) DeleteGuestHostRatings(guestId uint, hostId uint, ctx context.Context) (int64, error) {
	return r.DeleteGuestHostRatingsFn(guestId, hostId, ctx)
}

func (r *MockRepo) RateHost(hostRating *model.HostRating, ctx context.Context) (*model.HostRating, error) {
	return r.RateHostFn(hostRating, ctx)
}

func (r *MockRepo) FindGuestAccomodationRatings(guestId uint, accomodationId uint, ctx context.Context) (*[]model.AccomodationRating, error) {
	return &[]model.AccomodationRating{}, nil

}

func (r *MockRepo) DeleteGuestAccomodationRatings(guestId uint, accomodationId uint, ctx context.Context) (int64, error) {
	return 0, nil
}

func (r *MockRepo) RateAccomodation(accomodationRating *model.AccomodationRating, ctx context.Context) (*model.AccomodationRating, error) {
	return accomodationRating, nil
}
