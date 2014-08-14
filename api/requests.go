package api

import (
    "log"
    "github.com/franela/goreq"
)

// Get request to an API endpoint from dockbroker.
func Get(url string, result interface{}) interface{} {
    resp, err := goreq.Request{
        Method: "GET",
        Uri: url,
    }.Do()
    if err != nil { panic(err) }
    if resp.StatusCode != 200 {
        log.Fatal("Unexpected status code", resp.StatusCode)
    }

    resp.Body.FromJsonTo(result)
    return result
}

// Post request to an API endpoint from dockbroker.
func Post(url string, payload interface{}, result interface{}) interface{} {
    resp, err := goreq.Request{
        Method: "POST",
        Uri: url,
        Body: payload,
    }.Do()
    if err != nil { panic(err) }
    if resp.StatusCode != 200 {
        log.Fatal("Unexpected status code", resp.StatusCode)
    }

    resp.Body.FromJsonTo(result)
    return result
}

