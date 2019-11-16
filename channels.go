package gobttv

import (
	"fmt"

	"github.com/dankeroni/jsonapi"
)

// ChannelResponse json to struct
type ChannelResponse struct {
	ID           string   `json:"id"`
	Bots         []string `json:"bots"`
	Emotes       []Emote  `json:"channelEmotes"`
	SharedEmotes []Emote  `json:"sharedEmotes"`
}

func (bttvAPI *BTTVAPI) getChannel(channelID string, onSuccess func(ChannelResponse),
	onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	var response ChannelResponse
	onSuccessfulRequest := func() {
		for i := range response.Emotes {
			response.Emotes[i].populateURLs()
		}

		for i := range response.SharedEmotes {
			response.SharedEmotes[i].populateURLs()
		}

		onSuccess(response)
	}
	bttvAPI.Get("/cached/users/twitch/"+channelID, nil, &response, onSuccessfulRequest,
		onHTTPError, onInternalError)
}

// GetChannel request for GET https://api.betterttv.net/3/cached/users/twitch/:channelID
func (bttvAPI *BTTVAPI) GetChannel(channelID string) (channel ChannelResponse, err error) {
	c := make(chan struct{})

	onSuccessfulRequest := func(r ChannelResponse) {
		channel = r

		close(c)
	}

	onHTTPError := func(statusCode int, statusMessage, errorMessage string) {
		if statusCode == 404 {
			// User has not signed into BTTV - return an empty struct
			close(c)
			return
		}

		err = fmt.Errorf("HTTP Error occurred while getting channel of %s: %d(%s) %s", channelID, statusCode, statusMessage, errorMessage)
		close(c)
	}

	onInternalError := func(rErr error) {
		err = rErr
		close(c)
	}

	bttvAPI.getChannel(channelID, onSuccessfulRequest, onHTTPError, onInternalError)

	<-c

	return
}
