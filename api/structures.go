//Package api provides the structures for the JSON interface.
package api

import (
    "time"
)

// Job describes a job.
type Job struct {
    JobName string
    Submitter string
    EstTime time.Time
    MaxTime time.Time
}

// Offer holds the data about a broker's offer.
type Offer struct {
    Price float32
    EstCompletionTime time.Time
}

// BrokerInfo describes the broker.
type BrokerInfo struct {
    Name string
    Note string
}
