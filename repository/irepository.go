package repository

import (
	"context"

	"github.com/windbnb/review-service/model"
)

type IRepository interface {
	RateHost(hostRating *model.HostRating, ctx context.Context) (*model.HostRating, error)
	RateAccomodation(ctx context.Context) (*model.AccomodationRating, error)
}
