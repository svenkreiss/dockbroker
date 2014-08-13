package main

import (
	"fmt"
	"time"
	"net/http"
	"encoding/json"
)

func makeJsonHandler(fn func(r *http.Request) interface{}) http.HandlerFunc {
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


type Info struct {
	BrokerName string
	Path string
}

func infoHandler(r *http.Request) interface{} {
	return Info{"Paul", r.URL.Path}
}


type Offer struct {
	Price float64
	EstCompletionTime time.Time
}

func offerHandler(r *http.Request) interface{} {
	return Offer{20.50, time.Now()}
}


func main() {
	fmt.Printf("Starting dockbroker daemon on port 4027.\n")
    http.HandleFunc("/info/", makeJsonHandler(infoHandler))
    http.HandleFunc("/offer/", makeJsonHandler(offerHandler))
    http.ListenAndServe(":4027", nil)
}
