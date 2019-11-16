package gobttv

import (
	"testing"
)

var api = New()

func TestEnsureNaMExists(t *testing.T) {
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
