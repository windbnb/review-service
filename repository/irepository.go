package repository

import (
	"context"

	"github.com/windbnb/review-service/model"
)

type IRepository interface {
	FindGuestHostRatings(guestId uint, hostId uint, ctx context.Context) (*[]model.HostRating, error)
	DeleteGuestHostRatings(guestId uint, hostId uint, ctx context.Context) (int64, error)
	RateHost(hostRating *model.HostRating, ctx context.Context) (*model.HostRating, error)

	FindGuestAccomodationRatings(guestId uint, accomodationId uint, ctx context.Context) (*[]model.AccomodationRating, error)
	DeleteGuestAccomodationRatings(guestId uint, accomodationId uint, ctx context.Context) (int64, error)
	RateAccomodation(accomodationRating *model.AccomodationRating, ctx context.Context) (*model.AccomodationRating, error)
}
