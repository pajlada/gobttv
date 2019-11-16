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
    // 11148817 is the user ID for pajlada
    api.GetEmotes("11148817", onSuccess, onHTTPError, onInternalError)
}

func onSuccess(emotes gobttv.EmotesResponse) {
    fmt.Printf("%+v\n", emotes)
}

func onHTTPError(statusCode int, statusMessage, errorMessage string) {
    fmt.Println("statusCode:", statusCode)
    fmt.Println("statusMessage:", statusMessage)
    fmt.Println("errorMessage:", errorMessage)
}

func onInternalError(err error) {
    fmt.Println("internalError:", err)
}
```
