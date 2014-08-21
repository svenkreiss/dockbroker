// dockbroker is a daemon that manages a node for peer-to-peer computing.
package main

import (
	"io"
	"fmt"
	"net/http"
	"encoding/json"
	"code.google.com/p/go.net/websocket"

	"github.com/svenkreiss/dockbroker/api"
)

func makeJSONHandler(fn func(r *http.Request) interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		obj := fn(r)

		js, err := json.Marshal(obj)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func handleJSONPostRequest(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	if err != nil { panic(err) }
}

func createOffer(manifest *api.Job) api.Offer {
	price := 0.0
	price += 2.0 * manifest.EstDuration.Hours()

	completionTime := Queue.EstTime()
	completionTime.Add(manifest.EstDuration)

	return api.Offer{price, completionTime}
}


// Returns a struct for an info request.
func infoHandler(r *http.Request) interface{} {
	return api.Broker{"Paul", r.URL.Path}
}

// Returns a struct for an offer request.
func offerHandler(r *http.Request) interface{} {
	// read in the offer
	manifest := new(api.Job)
	handleJSONPostRequest(r, manifest)

	// create and send the offer for this request
	return createOffer(manifest)
}

// Returns a struct for a submitted job.
func submitHandler(r *http.Request) interface{} {
	manifest := new(api.Job)
	handleJSONPostRequest(r, manifest)

	fmt.Printf("open websocket connections: %v\n", wsConnections)

	return api.SubmittedJob{createOffer(manifest), Queue.NewJob(*manifest)}
}





var wsConnections = map[string]*websocket.Conn {}

// Handle WebSocket connections
func wsHandler(ws *websocket.Conn) {
	// add this connection to global list
	wsConnections["test"] = ws

	// handle messages
	var msg interface{}
	for {
		err := websocket.JSON.Receive(ws, &msg)
		if err != nil {
			if err == io.EOF { break }
			panic(err)
		}
		fmt.Printf("Received from server: %v\n", msg)
		err = websocket.JSON.Send(ws, msg)
		if err != nil { panic(err) }
	}
}






func main() {
	fmt.Printf("Starting dockbroker daemon on port 4027.\n")
    http.HandleFunc("/api/info/", makeJSONHandler(infoHandler))
    http.HandleFunc("/api/offer/", makeJSONHandler(offerHandler))
    http.HandleFunc("/api/submit/", makeJSONHandler(submitHandler))
    http.HandleFunc("/api/ws/", func (w http.ResponseWriter, req *http.Request) {
    	// this disables the origin check
        s := websocket.Server{Handler: websocket.Handler(wsHandler)}
        s.ServeHTTP(w, req)
    })
    http.ListenAndServe(":4027", nil)
}
