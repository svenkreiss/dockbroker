//Package api provides the structures for the JSON API.
package api

import (
    "time"
)

// Job describes a job.
type Job struct {
    Name string
    Submitter string
    MaxDuration time.Duration
    EstDuration time.Duration
}

// SubmittedJob describes a job that is in the queue now.
type SubmittedJob struct {
    Offer
    ID int
}

// Offer holds the data about a broker's offer.
type Offer struct {
    Price float64
    EstCompletionTime time.Time
}

// Broker describes the broker.
type Broker struct {
    Name string
    Note string
}
