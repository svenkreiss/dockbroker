// dockbroker is a daemon that manages a node for peer-to-peer computing.
package main

import (
	"fmt"
	"time"
	"net/http"
	"encoding/json"
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


// Info structure to be returned in the /info/ api.
type Info struct {
	BrokerName string
	Path string
}

// Returns a struct for an info request.
func infoHandler(r *http.Request) interface{} {
	return Info{"Paul", r.URL.Path}
}


// Offer structure to be returned in the /offer/ api.
type Offer struct {
	Price float64
	EstCompletionTime time.Time
}

// Returns a struct for an offer request.
func offerHandler(r *http.Request) interface{} {
	return Offer{20.50, time.Now()}
}


func main() {
    fmt.Printf("Queuing fake jobs.\n")
    Queue.NewJob("test job")
    Queue.NewJob("test job1")
    Queue.NewJob("test job2")
    Queue.NewJob("test job3")
    fmt.Printf("Printing fake jobs.\n")
    Queue.Print()

	fmt.Printf("Starting dockbroker daemon on port 4027.\n")
    http.HandleFunc("/info/", makeJSONHandler(infoHandler))
    http.HandleFunc("/offer/", makeJSONHandler(offerHandler))
    http.ListenAndServe(":4027", nil)
}
