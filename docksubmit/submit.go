// docksubmit is a command line tool to submit jobs to dockbroker.
package main

import (
	"io"
	"log"
	"time"
	"code.google.com/p/go.net/websocket"

	"github.com/svenkreiss/dockbroker/api"
)

func listenWs(ws *websocket.Conn) {
	for {
		// var msg interface{}
		// websocket.JSON.Receive(ws, msg)
		var msg interface{}
		err := websocket.JSON.Receive(ws, &msg)
		if err != nil {
			if err == io.EOF { break }
			panic(err)
		}
		log.Printf("Received ws message: %v\n", msg)
	}
}

func main() {
	log.Printf("Contacting broker on port 4027.\n")
	log.Printf("Info: %v", api.Get("http://localhost:4027/api/info/", new(api.Broker)))

	// opening ws connection
	ws, err := websocket.Dial("ws://localhost:4027/api/ws/", "", "http://localhost")
	if err != nil { panic(err) }
	// listening for ws message
	go listenWs(ws)

	// define the job
	manifest := api.Job{
		"test job 1",
		"Sven Kreiss <me@svenkreiss.com>",
		24 * time.Hour,
		12 * time.Hour,
	}

	websocket.JSON.Send(ws, manifest)

	// get an offer
	log.Printf("Offer: %v", api.Post("http://localhost:4027/api/offer/",
		&manifest, new(api.Offer)))

	// submit the job
	log.Printf("Submit: %v", api.Post("http://localhost:4027/api/submit/",
		&manifest, new(api.SubmittedJob)))
}
