package gobttv

import "github.com/dankeroni/jsonapi"

// ChannelEmoteData json to struct
type ChannelEmoteData struct {
	ID        string `json:"id"`
	Channel   string `json:"channel"`
	Code      string `json:"code"`
	ImageType string `json:"imageType"`
}

// ChannelResponse json to struct
type ChannelResponse struct {
	Status      int                `json:"status"`
	URLTemplate string             `json:"urlTemplate"`
	Bots        []string           `json:"bots"`
	Emotes      []ChannelEmoteData `json:"emotes"`
}

// GetChannel request for GET https://api.betterttv.net/channels/:channelName
func (bttvAPI *BTTVAPI) GetChannel(channelName string, onSuccess func(ChannelResponse),
	onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	var channel ChannelResponse
	onSuccessfulRequest := func() {
		onSuccess(channel)
	}
	bttvAPI.Get("/2/channels/"+channelName, nil, &channel, onSuccessfulRequest,
		onHTTPError, onInternalError)
}
