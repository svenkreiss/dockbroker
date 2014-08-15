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

    // define the job
    manifest := api.Job{
        "test job 1",
        "Sven Kreiss <me@svenkreiss.com>",
        24 * time.Hour,
        12 * time.Hour,
    }

    // get an offer
    log.Printf("Offer: %v", api.Post("http://localhost:4027/offer/",
        &manifest, new(api.Offer)))

    // submit the job
    log.Printf("Submit: %v", api.Post("http://localhost:4027/submit/",
        &manifest, new(api.SubmittedJob)))
}
