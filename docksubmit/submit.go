// docksubmit is a command line tool to submit jobs to dockbroker.
package main

import (
    "log"
    "time"

    "github.com/svenkreiss/dockbroker/api"
)

func main() {
    log.Printf("Contacting broker on port 4027.\n")
    log.Printf("Info: %v", api.Get("http://localhost:4027/info/", new(api.Broker)))
    log.Printf("Offer: %v", api.Post("http://localhost:4027/offer/",
        &api.Job{
            "test job 1",
            "Sven Kreiss <me@svenkreiss.com>",
            24 * time.Hour,
            12 * time.Hour,
        }, new(api.Offer)))
}
