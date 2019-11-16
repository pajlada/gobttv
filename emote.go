package gobttv

const emoteURLPrefix = "https://cdn.betterttv.net/emote/"

// Emote is the lowest common description of the BetterTTV emote
type Emote struct {
	// Fetched straight from the API
	ID        string `json:"id"`
	Code      string `json:"code"`
	ImageType string `json:"imageType"`

	// Filled in by us with populateURLs
	URLs URLs
}

func (e *Emote) populateURLs() {
	var emoteURL = emoteURLPrefix + e.ID
	e.URLs.X1 = emoteURL + "/1x"
	e.URLs.X2 = emoteURL + "/2x"
	e.URLs.X4 = emoteURL + "/3x"
}
