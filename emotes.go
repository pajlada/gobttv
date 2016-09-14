package gobttv

import "github.com/dankeroni/jsonapi"

// GlobalEmoteData json to struct
type GlobalEmoteData struct {
	URL         string      `json:"url"`
	Width       int         `json:"width"`
	Height      int         `json:"height"`
	ImageType   string      `json:"imageType"`
	Regex       string      `json:"regex"`
	Channel     interface{} `json:"channel"`
	EmoticonSet string      `json:"emoticon_set,omitempty"`
}

// EmotesResponse json to struct
type EmotesResponse struct {
	Status int               `json:"status"`
	Emotes []GlobalEmoteData `json:"emotes"`
}

// GetEmotes request for GET https://api.betterttv.net/emotes
func (bttvAPI *BTTVAPI) GetEmotes(onSuccess func(EmotesResponse),
	onHTTPError jsonapi.HTTPErrorCallback, onInternalError jsonapi.InternalErrorCallback) {
	var emotes EmotesResponse
	onSuccessfulRequest := func() {
		onSuccess(emotes)
	}
	bttvAPI.Get("/emotes", nil, &emotes, onSuccessfulRequest,
		onHTTPError, onInternalError)
}
