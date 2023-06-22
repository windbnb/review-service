package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type HostRating struct {
	ID      primitive.ObjectID `bson:"_id"`
	GuestId uint               `bson:"guestId"`
	HostId  uint               `bson:"hostId"`
	Raiting uint               `bson:"rating"`
}

type AccomodationRating struct {
	ID             primitive.ObjectID `bson:"_id"`
	GuestId        uint               `bson:"guestId"`
	AccomodationId uint               `bson:"accomodationId"`
	Raiting        uint               `bson:"rating"`
}
