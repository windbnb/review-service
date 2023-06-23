package model

type HostRatingRequest struct {
	GuestId uint
	HostId  uint
	Rating  uint
}

type AccomodationRatingRequest struct {
	GuestId        uint
	AccomodationId uint
	Rating         uint
}

type ErrorResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}
