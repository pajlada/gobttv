package gobttv

import (
	"testing"
)

var api = New()

func TestEnsureNaMExists(t *testing.T) {
	NaM := Emote{
		ID:        "566ca06065dbbdab32ec054e",
		Code:      "NaM",
		ImageType: "png",

		URLs: URLs{
			X1: "https://cdn.betterttv.net/emote/566ca06065dbbdab32ec054e/1x",
			X2: "https://cdn.betterttv.net/emote/566ca06065dbbdab32ec054e/2x",
			X4: "https://cdn.betterttv.net/emote/566ca06065dbbdab32ec054e/3x",
		},
	}

	foundNaM := false

	c := make(chan struct{})
	onResponse := func(emotes EmotesResponse) {
		for _, emote := range emotes {
			if emote == NaM {
				foundNaM = true
				break
			}
		}

		close(c)
	}

	api.GetEmotes(onResponse, onHTTPError(t), onInternalError(t))

	<-c
	if !foundNaM {
		t.Fatalf("Did not find NaM in set of global emotes, this might be due to the API being changed")
	}
}
