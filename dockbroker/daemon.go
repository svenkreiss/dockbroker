// dockbroker is a daemon that manages a node for peer-to-peer computing.
package main

import (
	"fmt"
	"net/http"
	"encoding/json"

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
	manifest := new(api.Job)
	handleJSONPostRequest(r, manifest)

	return createOffer(manifest)
}

// Returns a struct for a submitted job.
func submitHandler(r *http.Request) interface{} {
	manifest := new(api.Job)
	handleJSONPostRequest(r, manifest)

	return api.SubmittedJob{createOffer(manifest), Queue.NewJob(*manifest)}
}


func main() {
	fmt.Printf("Starting dockbroker daemon on port 4027.\n")
    http.HandleFunc("/info/", makeJSONHandler(infoHandler))
    http.HandleFunc("/offer/", makeJSONHandler(offerHandler))
    http.HandleFunc("/submit/", makeJSONHandler(submitHandler))
    http.ListenAndServe(":4027", nil)
}
