package main

import (
	"flag"
	"log"
	"math/rand"
	"sync"
	"time"
)

var (
	types   = []string{"emai", "deactivation", "activation", "transaction", "customer_renew", "order_processed"}
	workers = 0
)

func init() {
	flag.IntVar(&workers, "workers", 10, "Number of workers to use")
}

func getType() string {
	return types[rand.Int()%len(types)]
}

// main entry point for the application
func main() {
	// parse the flags
	flag.Parse()

	//////////
	// Demo of Worker Processing
	//////////

	// create a channel with a 10,000 Job buffer
	jobsChannel := make(chan *Job, 10000)

	// start the job processor
	go startJobProcessor(jobsChannel)

	createJobs(jobsChannel)
}

type Job struct {
	Type  string
	Sleep time.Duration
}

// makeJob creates a new job with a random sleep time between 10 ms and 4000ms
func makeJob() *Job {
	return &Job{
		Type:  getType(),
		Sleep: time.Duration(rand.Int()%100+10) * time.Millisecond,
	}
}

func startJobProcessor(jobs <-chan *Job) {
	log.Printf("[INFO] starting %d workers\n", workers)
	wait := sync.WaitGroup{}
	// notify the sync group we need to wait for 10 goroutines
	wait.Add(workers)

	// start 10 works
	for i := 0; i < workers; i++ {
		go func(workerID int) {
			// start the worker
			startWorker(workerID, jobs)
			wait.Done()
		}(i)
	}

	wait.Wait()
}

func createJobs(jobs chan<- *Job) {
	for {
		// create a random job
		job := makeJob()
		// send the job down the channel
		jobs <- job
		// don't pile up too quickly
		time.Sleep(5 * time.Millisecond)
	}
}

// creates a worker that pulls jobs from the job channel
func startWorker(workerID int, jobs <-chan *Job) {
	for {
		select {
		// read from the job channel
		case job := <-jobs:
			startTime := time.Now()
			// fake processing the request
			time.Sleep(job.Sleep)
			log.Printf("[%d][%s] Processed job in %0.3f seconds", workerID, job.Type, time.Now().Sub(startTime).Seconds())
		}
	}
}
