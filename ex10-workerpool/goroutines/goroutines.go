package goroutines

import "sync"

func Run(poolSize int) {
	var wg sync.WaitGroup
	maxWorkers := 100

	job := make(chan float64, maxWorkers)
	reader(job)
	close(job)
	creatWorker(poolSize, job, wg)
	wg.Wait()
}
