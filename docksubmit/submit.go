// docksubmit is a command line tool to submit jobs to dockbroker.
package main

import (
    "log"
    "net/http"
    "io/ioutil"
)

// Simple request to an API endpoint from dockbroker.
func brokerGet(url string) []byte {
    resp, err := http.Get(url)
    if err != nil { panic(err) }
    if resp.StatusCode != 200 {
        log.Fatal("Unexpected status code", resp.StatusCode)
    }
    body, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    if err != nil { panic(err) }
    return body
}

func main() {
    log.Printf("Contacting broker on port 4027.\n")
    log.Printf("Info: %s", brokerGet("http://localhost:4027/info/"))
    log.Printf("Offer: %s", brokerGet("http://localhost:4027/offer/"))
}
