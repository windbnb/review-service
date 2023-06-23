package repository

import (
	"context"
	"time"

	"github.com/windbnb/review-service/model"
	"github.com/windbnb/review-service/tracer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Db *mongo.Database
}

func (r *Repository) FindGuestHostRatings(guestId, hostId uint, ctx context.Context) (*[]model.HostRating, error) {
	span := tracer.StartSpanFromContext(ctx, "findGuestHostRatingsRepository")
	defer span.Finish()

	hostRatings := []model.HostRating{}
	dbCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	filter := bson.D{
		{"guestId", guestId},
		{"hostId", hostId},
	}

	cursor, err := r.Db.Collection("host-ratings").Find(dbCtx, filter)
	if err != nil {
		tracer.LogError(span, err)
		return nil, err
	}

	for cursor.Next(dbCtx) {
		var hostRating model.HostRating
		err := cursor.Decode(&hostRating)
		if err != nil {
			tracer.LogError(span, err)
			continue
		}

		hostRatings = append(hostRatings, hostRating)
	}

	return &hostRatings, nil

}

func (r *Repository) DeleteGuestHostRatings(guestId, hostId uint, ctx context.Context) (int64, error) {
	span := tracer.StartSpanFromContext(ctx, "findGuestHostRatingsRepository")
	defer span.Finish()

	dbCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	filter := bson.D{
		{"guestId", guestId},
		{"hostId", hostId},
	}

	data, err := r.Db.Collection("host-ratings").DeleteMany(dbCtx, filter)
	if err != nil {
		tracer.LogError(span, err)
		return 0, err
	}

	return data.DeletedCount, nil
}

func (r *Repository) RateHost(hostRating *model.HostRating, ctx context.Context) (*model.HostRating, error) {
	span := tracer.StartSpanFromContext(ctx, "submitHostRating")
	defer span.Finish()

	dbCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	hostRating.ID = primitive.NewObjectID()
	_, err := r.Db.Collection("host-ratings").InsertOne(dbCtx, &hostRating)
	if err != nil {
		tracer.LogError(span, err)
		return nil, err
	}

	return hostRating, nil
}

func (r *Repository) FindAllHostRatings(hostId uint, ctx context.Context) (*[]model.HostRating, error) {
	span := tracer.StartSpanFromContext(ctx, "findAllHostRatingsRepository")
	defer span.Finish()

	hostRatings := []model.HostRating{}
	dbCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	filter := bson.D{
		{"hostId", hostId},
	}

	cursor, err := r.Db.Collection("host-ratings").Find(dbCtx, filter)
	if err != nil {
		tracer.LogError(span, err)
		return nil, err
	}

	for cursor.Next(dbCtx) {
		var hostRating model.HostRating
		err := cursor.Decode(&hostRating)
		if err != nil {
			tracer.LogError(span, err)
			continue
		}

		hostRatings = append(hostRatings, hostRating)
	}

	return &hostRatings, nil

}

func (r *Repository) FindGuestAccomodationRatings(guestId, accomodationId uint, ctx context.Context) (*[]model.AccomodationRating, error) {
	span := tracer.StartSpanFromContext(ctx, "findGuestAccomodationRatingsRepository")
	defer span.Finish()

	accomodationsRatings := []model.AccomodationRating{}
	dbCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	filter := bson.D{
		{"guestId", guestId},
		{"accomodationId", accomodationId},
	}

	cursor, err := r.Db.Collection("accomodation-ratings").Find(dbCtx, filter)
	if err != nil {
		tracer.LogError(span, err)
		return nil, err
	}

	for cursor.Next(dbCtx) {
		var accomodationsRating model.AccomodationRating
		err := cursor.Decode(&accomodationsRating)
		if err != nil {
			tracer.LogError(span, err)
			continue
		}

		accomodationsRatings = append(accomodationsRatings, accomodationsRating)
	}

	return &accomodationsRatings, nil

}

func (r *Repository) DeleteGuestAccomodationRatings(guestId, accomodationId uint, ctx context.Context) (int64, error) {
	span := tracer.StartSpanFromContext(ctx, "deleteGuestAccomodationRatingsRepository")
	defer span.Finish()

	dbCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	filter := bson.D{
		{"guestId", guestId},
		{"accomodationId", accomodationId},
	}

	data, err := r.Db.Collection("accomodation-ratings").DeleteMany(dbCtx, filter)
	if err != nil {
		tracer.LogError(span, err)
		return 0, err
	}

	return data.DeletedCount, nil
}

func (r *Repository) RateAccomodation(accomodationRating *model.AccomodationRating, ctx context.Context) (*model.AccomodationRating, error) {
	span := tracer.StartSpanFromContext(ctx, "submitAccomodationRating")
	defer span.Finish()

	dbCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	accomodationRating.ID = primitive.NewObjectID()
	_, err := r.Db.Collection("accomodation-ratings").InsertOne(dbCtx, &accomodationRating)
	if err != nil {
		tracer.LogError(span, err)
		return nil, err
	}

	return accomodationRating, nil
}

func (r *Repository) FindAllAccomodationRatings(accomodationId uint, ctx context.Context) (*[]model.AccomodationRating, error) {
	span := tracer.StartSpanFromContext(ctx, "findAllAccomodationRatingsRepository")
	defer span.Finish()

	accomodationsRatings := []model.AccomodationRating{}
	dbCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	filter := bson.D{
		{"accomodationId", accomodationId},
	}

	cursor, err := r.Db.Collection("accomodation-ratings").Find(dbCtx, filter)
	if err != nil {
		tracer.LogError(span, err)
		return nil, err
	}

	for cursor.Next(dbCtx) {
		var accomodationsRating model.AccomodationRating
		err := cursor.Decode(&accomodationsRating)
		if err != nil {
			tracer.LogError(span, err)
			continue
		}

		accomodationsRatings = append(accomodationsRatings, accomodationsRating)
	}

	return &accomodationsRatings, nil

}
