// Package workerpool implements a simple goroutine pool.
package workerpool

import (
	"sync"
)

// Job is an entity that works on a specific task (the Process function).
type Job interface {
	Process() JobResult
}

// JobResult is a result of the task performed by Job.
type JobResult interface {
	GetJob() Job
}

// WorkerPool is an implementation of a Pool of Goroutines.
type WorkerPool struct {
	nWorkers       int
	pendingJobs    chan Job
	pendingAddJobs *sync.WaitGroup
	finishedJobs   chan JobResult
	workersActive  *sync.WaitGroup
}

// New generates a WorkerPool struct and runs "nWorkers" workers.
func New(nWorkers int) *WorkerPool {
	pool := &WorkerPool{
		nWorkers:       nWorkers,
		pendingJobs:    make(chan Job),
		pendingAddJobs: &sync.WaitGroup{},
		finishedJobs:   make(chan JobResult),
		workersActive:  &sync.WaitGroup{},
	}

	return pool
}

// Run initiates the Worker Pool.
func (pool *WorkerPool) Run() {
	// Create a goroutine for each worker:
	for i := 0; i < pool.nWorkers; i++ {
		pool.workersActive.Add(1)
		go workerRoutine(pool)
	}

	go waitForWorkersRoutine(pool)
}

// GetResultsChannel returns the channel from which job results can we collected.
func (pool *WorkerPool) GetResultsChannel() chan JobResult {
	return pool.finishedJobs
}

// AddJob adds a job to the pool of workers.
func (pool *WorkerPool) AddJob(job Job) {
	pool.pendingAddJobs.Add(1)
	go func() {
		pool.pendingJobs <- job
		pool.pendingAddJobs.Done()
	}()
}

// EndJobs tells the Worker Pool that there are no more jobs incoming
// This internally closes the channel for incoming jobs.
func (pool *WorkerPool) EndJobs() {
	// This needs to be done only after all the jobs have been added to the pendingJobs channel:
	go func() {
		pool.pendingAddJobs.Wait()
		close(pool.pendingJobs)
	}()
}

// workerRoutine corresponds to the routine in which a worker runs until it is done.
func workerRoutine(pool *WorkerPool) {
	// While there are jobs to process:
	for job := range pool.pendingJobs {
		result := job.Process()
		pool.finishedJobs <- result
	}

	// Mark this worker as finished:
	pool.workersActive.Done()
}

// waitForWorkersRoutine corresponds to a routine that waits until all workers finish.
func waitForWorkersRoutine(pool *WorkerPool) {
	pool.workersActive.Wait()
	close(pool.finishedJobs)
}
