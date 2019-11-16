package gobttv

import "github.com/dankeroni/jsonapi"

type test struct {
	asd string
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
