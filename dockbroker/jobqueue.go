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
    fmt.Printf("job: name=%s, id=%d, estdur=%v, started=%v, finished=%v\n",
               j.Manifest.Name, j.id, j.Manifest.EstDuration, j.Started, j.Finished)
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
    jq.Print()
}

// Print an overview of the current queue.
func (jq *JobQueue) Print() {
    fmt.Println("--- current queue ---")
    for j := jq.l.Front(); j != nil; j = j.Next() {
        j.Value.(QueueElement).Print()
    }
}

// NewJob creates a new Job and enqueues it.
func (jq *JobQueue) NewJob(manifest api.Job) int {
    jq.lastID++

    job := new(Job)
    job.Manifest = manifest
    job.id = jq.lastID
    jq.Enqueue(job)

    return jq.lastID
}

// EstTime estimates the time when this queue will be completed.
func (jq *JobQueue) EstTime() time.Time {
    r := time.Now()
    for j := jq.l.Front(); j != nil; j = j.Next() {
        job := j.Value.(*Job)

        if job.Started == nil && job.Finished == nil {
            // job hasn't started yet
            r = r.Add(job.Manifest.EstDuration)
        } else if job.Started != nil && job.Finished == nil {
            // job has started and is running

            // is it before the estimated finish time?
            if time.Since(*job.Started) < job.Manifest.EstDuration {
                r = r.Add(job.Manifest.EstDuration - time.Since(*job.Started))
            }
        }
    }
    return r
}

// the job queue for this broker
var Queue = JobQueue{list.New(), 0}

