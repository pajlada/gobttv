package gobttv

import "testing"

const (
	userIDNotLoggedIn      = "14027"
	userIDLoggedInNoEmotes = "117166826"
	userIDPajlada          = "11148817"
)

var WAYTOODANK = Emote{
	ID:        "5ad22a7096065b6c6bddf7f3",
	Code:      "WAYTOODANK",
	ImageType: "gif",

	URLs: URLs{
		X1: "https://cdn.betterttv.net/emote/5ad22a7096065b6c6bddf7f3/1x",
		X2: "https://cdn.betterttv.net/emote/5ad22a7096065b6c6bddf7f3/2x",
		X4: "https://cdn.betterttv.net/emote/5ad22a7096065b6c6bddf7f3/3x",
	},
}

var NaM = Emote{
	ID:        "566ca06065dbbdab32ec054e",
	Code:      "NaM",
	ImageType: "png",

	URLs: URLs{
		X1: "https://cdn.betterttv.net/emote/566ca06065dbbdab32ec054e/1x",
		X2: "https://cdn.betterttv.net/emote/566ca06065dbbdab32ec054e/2x",
		X4: "https://cdn.betterttv.net/emote/566ca06065dbbdab32ec054e/3x",
	},
}

func onHTTPError(t *testing.T) func(statusCode int, statusMessage, errorMessage string) {
	return func(statusCode int, statusMessage, errorMessage string) {
		t.Fatalf("HTTP error %d (%s): %s", statusCode, statusMessage, errorMessage)
	}
}

func onInternalError(t *testing.T) func(err error) {
	return func(err error) {
		t.Fatal(err)
	}
}
