package gobttv

import "github.com/dankeroni/jsonapi"

// ChannelResponse json to struct
type ChannelResponse struct {
	ID            string   `json:"id"`
	Bots          []string `json:"bots"`
	ChannelEmotes []Emote  `json:"channelEmotes"`
	SharedEmotes  []Emote  `json:"sharedEmotes"`
}

// GetChannel request for GET https://api.betterttv.net/3/cached/users/twitch/:channelID
func (bttvAPI *BTTVAPI) GetChannel(channelID string, onSuccess func(ChannelResponse),
	onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	var response ChannelResponse
	onSuccessfulRequest := func() {
		for i := range response.ChannelEmotes {
			response.ChannelEmotes[i].populateURLs()
		}

		for i := range response.SharedEmotes {
			response.SharedEmotes[i].populateURLs()
		}

		onSuccess(response)
	}
	bttvAPI.Get("/cached/users/twitch/"+channelID, nil, &response, onSuccessfulRequest,
		onHTTPError, onInternalError)
}
