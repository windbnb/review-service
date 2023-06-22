package repository

import (
	"context"
	"time"

	"github.com/windbnb/review-service/model"
	"github.com/windbnb/review-service/tracer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Db *mongo.Database
}

func (r *Repository) RateHost(hostRating *model.HostRating, ctx context.Context) (*model.HostRating, error) {
	span := tracer.StartSpanFromContext(ctx, "submitHostRating")
	defer span.Finish()

	dbCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	filter := bson.D{
		// {"accommodationID", accomodationId},
		// {"status", model.ACCEPTED},
	}

	cursor, err := r.Db.Collection("host_ratings").Find(dbCtx, filter)
	if err != nil {
		tracer.LogError(span, err)
		return nil, nil
	}

	cursor.Next(dbCtx)
	var hostRatingData model.HostRating
	err = cursor.Decode(&hostRatingData)
	if err != nil {
		tracer.LogError(span, err)
	}

	// for cursor.Next(dbCtx) {
	// 	var hostRating model.HostRating
	// 	err := cursor.Decode(&hostRating)
	// 	if err != nil {
	// 		tracer.LogError(span, err)
	// 		continue
	// 	}

	// 	hostRatings = append(hostRatings, hostRating)
	// }

	return &hostRatingData, nil
}

func (r *Repository) RateAccomodation(ctx context.Context) (*model.AccomodationRating, error) {
	return &model.AccomodationRating{}, nil
}
