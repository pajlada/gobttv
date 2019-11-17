package gobttv

import (
	"fmt"

	"github.com/dankeroni/jsonapi"
)

// EmotesResponse json to struct
type EmotesResponse []Emote

// getEmotes request for GET https://api.betterttv.net/3/cached/emotes/global
func (bttvAPI *BTTVAPI) getEmotes(onSuccess func(EmotesResponse),
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

// GetEmotes request for GET https://api.betterttv.net/3/cached/emotes/global
func (bttvAPI *BTTVAPI) GetEmotes() (emotes EmotesResponse, err error) {
	c := make(chan struct{})
	onSuccessfulRequest := func(r EmotesResponse) {
		for i := range r {
			r[i].populateURLs()
		}

		emotes = r

		close(c)
	}

	onHTTPError := func(statusCode int, statusMessage, errorMessage string) {
		err = fmt.Errorf("HTTP Error occurred while getting global emotes: %d(%s) %s", statusCode, statusMessage, errorMessage)
		close(c)
	}

	onInternalError := func(rErr error) {
		err = rErr
		close(c)
	}

	bttvAPI.getEmotes(onSuccessfulRequest,
		onHTTPError, onInternalError)

	<-c

	return
}
