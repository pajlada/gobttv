package gobttv

import (
	"testing"
)

func TestEnsureNaMExistsInternal(t *testing.T) {
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

	api.getEmotes(onResponse, onHTTPError(t), onInternalError(t))

	<-c
	if !foundNaM {
		t.Fatal("Did not find NaM in set of global emotes, this might be due to the API being changed")
	}
}

func TestEnsureNaMExists(t *testing.T) {
	emotes, err := api.GetEmotes()
	if err != nil {
		t.Fatal(err)
	}

	for _, emote := range emotes {
		if emote == NaM {
			return
		}
	}

	t.Fatal("Did not find NaM in set of global emotes, this might be due to the API being changed")
}
