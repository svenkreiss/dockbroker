// Package jobqueue provides the book-keeping for the dockbroker jobs.
package main

import (
    "fmt"
    "container/list"
)

type Job struct {
    Name string
    id int
}

// the job queue for this broker
var queue = list.New()
var lastId = 0

// add a job to the queue
func Enqueue(job Job) {
    queue.PushBack(job)
}

// create a new Job and Enqueue() it
func NewJob(name string) {
    lastId += 1
    job := Job{name, lastId}
    Enqueue(job)
}

// print an overview of the current queue
func PrintQueue() {
    for j := queue.Front(); j != nil; j = j.Next() {
        fmt.Printf("job: %v\n", j.Value)
    }
}
