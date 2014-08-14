package main

import (
    "fmt"
    "container/list"
)



type Job struct {
    Name string
    id int
}

func (j *Job) Print() {
    fmt.Printf("job: name=%s, id=%d\n", j.Name, j.id)
}



type JobQueue struct {
    l *list.List
    lastId int
}

type QueueElement interface {
    Print()
}

// add a job to the queue
func (jq *JobQueue) Enqueue(qe QueueElement) {
    jq.l.PushBack(qe)
}

// print an overview of the current queue
func (jq *JobQueue) Print() {
    for j := jq.l.Front(); j != nil; j = j.Next() {
        j.Value.(QueueElement).Print()
    }
}

// create a new Job and Enqueue() it
func (jq *JobQueue) NewJob(name string) {
    jq.lastId++
    job := Job{name, jq.lastId}
    jq.Enqueue(&job)
}

// the job queue for this broker
var Queue = JobQueue{list.New(), 0}

