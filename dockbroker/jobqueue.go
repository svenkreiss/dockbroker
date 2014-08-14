package main

import (
    "fmt"
    "time"
    "container/list"

    "github.com/svenkreiss/dockbroker/api"
)


// Job description
type Job struct {
    Manifest api.Job
    id int

    Started *time.Time
    Finished *time.Time
}

// Print Job information.
func (j *Job) Print() {
    fmt.Printf("job: name=%s, id=%d, started=%v, finished=%v\n",
               j.Manifest.Name, j.id, j.Started, j.Finished)
}


// JobQueue is the dockbroker queue of jobs
type JobQueue struct {
    l *list.List
    lastID int
}

// QueueElement is a more generic interface (than the specific Job structure)
// for the JobQueue.
type QueueElement interface {
    Print()
}

// Enqueue adds a job to the queue.
func (jq *JobQueue) Enqueue(qe QueueElement) {
    jq.l.PushBack(qe)
}

// Print an overview of the current queue.
func (jq *JobQueue) Print() {
    for j := jq.l.Front(); j != nil; j = j.Next() {
        j.Value.(QueueElement).Print()
    }
}

// NewJob creates a new Job and enqueues it.
func (jq *JobQueue) NewJob(manifest api.Job) {
    jq.lastID++
    job := new(Job)
    job.Manifest = manifest
    job.id = jq.lastID
    jq.Enqueue(job)
}

// the job queue for this broker
var Queue = JobQueue{list.New(), 0}

