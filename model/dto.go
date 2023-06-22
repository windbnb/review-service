package model

type HostRatingRequest struct {
	GuestId uint
	HostId  uint
	Raiting uint
}

type AccomodationRatingRequest struct {
	GuestId        uint
	AccomodationId uint
	Raiting        uint
}

type ErrorResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}
