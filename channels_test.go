package gobttv

import "testing"

func TestFindChannelEmoteWAYTOODANK(t *testing.T) {
	e := Emote{
		ID:        "5ad22a7096065b6c6bddf7f3",
		Code:      "WAYTOODANK",
		ImageType: "gif",

		URLs: URLs{
			X1: "https://cdn.betterttv.net/emote/5ad22a7096065b6c6bddf7f3/1x",
			X2: "https://cdn.betterttv.net/emote/5ad22a7096065b6c6bddf7f3/2x",
			X4: "https://cdn.betterttv.net/emote/5ad22a7096065b6c6bddf7f3/3x",
		},
	}

	found := false

	c := make(chan struct{})
	onResponse := func(response ChannelResponse) {
		for _, emote := range response.ChannelEmotes {
			if emote == e {
				found = true
				break
			}
		}

		close(c)
	}

	api.GetChannel("11148817", onResponse, onHTTPError(t), onInternalError(t))

	<-c
	if !found {
		t.Fatalf("Did not find WAYTOODANK in set of channel emotes, this might be due to the API being changed")
	}

}
