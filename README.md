# gobttv

Uses [jsonapi](https://github.com/dankeroni/jsonapi) by [dankeroni](https://github.com/dankeroni) for all net code.

Inspired by [gotwitch](https://github.com/dankeroni/gotwitch)

Uses BetterTTV's v3 api

## Example for getting a Channel object
```go
package main

import (
    "fmt"

    "github.com/pajlada/gobttv"
)

var api = gobttv.New()

func main() {
    // Get channel information for user "pajlada" (user id 11148817)
    channel, err := api.GetChannel("11148817")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Channel emotes:", channel.Emotes)
    fmt.Println("Channel shared emotes:", channel.SharedEmotes)

    // Get global emotes
    emotes, err := api.GetEmotes()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Global emotes:", emotes)
}
```
