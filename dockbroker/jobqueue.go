package main

import (
    "fmt"
    "container/list"
)


// Job description
type Job struct {
    Name string
    id int
}

// Print Job information.
func (j *Job) Print() {
    fmt.Printf("job: name=%s, id=%d\n", j.Name, j.id)
}


// JobQueue is the dockbroker queue of jobs
type JobQueue struct {
    l *list.List
    lastID int
}

// QueueElement is a more generic interface (than specific the Job structure)
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
func (jq *JobQueue) NewJob(name string) {
    jq.lastID++
    job := Job{name, jq.lastID}
    jq.Enqueue(&job)
}

// the job queue for this broker
var Queue = JobQueue{list.New(), 0}

