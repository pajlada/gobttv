package gobttv

import (
	"testing"
)

func assertNoChannelResponse(t *testing.T) func(channel *ChannelResponse) {
	return func(channel *ChannelResponse) {
		t.Fatal("Got channel response when none was expected")
	}
}

func TestFindChannelEmoteWAYTOODANK(t *testing.T) {
	channel, err := api.GetChannel(userIDPajlada)
	if err != nil {
		t.Fatal(err)
	}

	for _, emote := range channel.Emotes {
		if emote == WAYTOODANK {
			return
		}
	}

	t.Fatal("Did not find WAYTOODANK in set of channel emotes, this might be due to the API being changed")
}

func TestChannelHaveNotSignedInInternal(t *testing.T) {
	onHTTPError := func(statusCode int, statusMessage, errorMessage string) {
		if statusCode != 404 {
			t.Fatalf("Unhandled HTTP error %d (%s): %s", statusCode, statusMessage, errorMessage)
		}
	}

	api.getChannel(userIDNotLoggedIn, assertNoChannelResponse(t), onHTTPError, onInternalError(t))
}

func TestChannelHaveNotSignedIn(t *testing.T) {
	channel, err := api.GetChannel(userIDNotLoggedIn)
	if err != nil {
		t.Fatal(err)
	}

	if len(channel.Emotes) != 0 || len(channel.SharedEmotes) != 0 {
		t.Fatal("There should not be any emotes in this response")
	}
}

func TestChannelHaveSignedInNoEmotesInternal(t *testing.T) {
	c := make(chan struct{})
	onResponse := func(channel *ChannelResponse) {
		if len(channel.Emotes) != 0 || len(channel.SharedEmotes) != 0 {
			t.Fatal("There should not be any emotes in this response")
		}
		close(c)
	}

	api.getChannel(userIDLoggedInNoEmotes, onResponse, onHTTPError(t), onInternalError(t))

	<-c
}

func TestChannelHaveSignedInNoEmotes(t *testing.T) {
	channel, err := api.GetChannel(userIDLoggedInNoEmotes)
	if err != nil {
		t.Fatal(err)
	}

	if len(channel.Emotes) != 0 || len(channel.SharedEmotes) != 0 {
		t.Fatal("There should not be any emotes in this response")
	}
}
