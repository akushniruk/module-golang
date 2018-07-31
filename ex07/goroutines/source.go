package goroutines

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

var count int = 0

type Worker struct {
	ID int
}

func creatWorker(pool int, jobs <-chan float64, wg sync.WaitGroup) {
	worker := Worker{}
	for worker.ID = 1; worker.ID <= pool; worker.ID++ {
		wg.Add(1)
		go startJob(worker.ID, jobs, &wg)
		time.Sleep(time.Millisecond)
	}
}

func reader(jobs chan float64) {
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		file := string(scan.Bytes())
		parse, err := strconv.ParseFloat(file, 64)
		checkErr(err)
		jobs <- parse
	}
}

func start(id int) {
	fmt.Printf("worker: %d\n", id)
	count++
	//fmt.Printf("NubmerOfActiveGoroutines: %d\n", count)
}

func stop(id int) {
	fmt.Printf("worker:%d stopping\n", id)
	count--
	//fmt.Printf("NumberOfActiveGorutines: %d\n", count)
}

func startJob(id int, jobs <-chan float64, wg *sync.WaitGroup) {
	for job := range jobs {
		start(id)
		fmt.Printf("worker:%d sleep:%.1f\n", id, job)
		time.Sleep(time.Second * 2)
		stop(id)
		wg.Done()
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("Massange", err)
	}
}

//flagi ub
//addErrorChecker
