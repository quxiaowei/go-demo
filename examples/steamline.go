package examples

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func startWork(worker *Worker, jobPool <-chan Job, wg *sync.WaitGroup) {
	for job := range jobPool {
		// take a rest
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		worker.job = &job
		time.Sleep(time.Duration(job.load) * time.Second)
		worker.job = nil
		wg.Done()
	}
}

func Steamline(workerCount int, jobCount int) {

	var wg sync.WaitGroup
	jobPool := make(chan Job, jobCount)
	workerList := GetWorkers(workerCount)

	fmt.Println(len(workerList), "WORKERs,", jobCount, "JOBs")
	fmt.Println("START")

	// Monitor Function
	go func() {
		var jobName string

		for i := range workerList {
			fmt.Printf("|%s\t", workerList[i].name)
		}
		fmt.Println("|")

		ticker := time.NewTicker(500 * time.Millisecond)
		for range ticker.C {
			for _, worker := range workerList {
				if worker.job != nil {
					jobName = strconv.Itoa(worker.job.id)
				} else {
					jobName = "-"
				}
				fmt.Printf("| %s\t", jobName)
			}
			fmt.Println("|")
		}
	}()

	// init jobs
	for i := 0; i < jobCount; i++ {
		jobPool <- Job{id: i + 1, load: rand.Intn(9) + 1}
		wg.Add(1)
	}

	// workers start
	for i := range workerList {
		go startWork(&workerList[i], jobPool, &wg)
	}

	// waiting
	wg.Wait()
	fmt.Println("FINISH")
}
