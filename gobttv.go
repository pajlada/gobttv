package gobttv

import "github.com/dankeroni/jsonapi"

type errorResponse struct {
	Error   string `json:"error"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// BTTVAPI struct
type BTTVAPI struct {
	jsonapi.JSONAPI
}

// New instantiates a new BTTVAPI object
func New() *BTTVAPI {
	return &BTTVAPI{
		JSONAPI: jsonapi.JSONAPI{
			BaseURL: "https://api.betterttv.net/3",
		},
	}
}
