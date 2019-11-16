package gobttv

import (
	"github.com/dankeroni/jsonapi"
)

// EmotesResponse json to struct
type EmotesResponse []Emote

// GetEmotes request for GET https://api.betterttv.net/3/cached/emotes/global
func (bttvAPI *BTTVAPI) GetEmotes(onSuccess func(EmotesResponse),
	onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	var emotes EmotesResponse
	onSuccessfulRequest := func() {
		for i := range emotes {
			emotes[i].populateURLs()
		}

		onSuccess(emotes)
	}

	bttvAPI.Get("/cached/emotes/global", nil, &emotes, onSuccessfulRequest,
		onHTTPError, onInternalError)
}
