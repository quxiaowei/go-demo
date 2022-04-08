package examples

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func _worker(worker *Worker, jobs <-chan int, wg *sync.WaitGroup) {
	for j := range jobs {
		// take a rest
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		jId := j
		worker.jobId = jId
		time.Sleep(time.Duration((rand.Intn(6) + 1)) * time.Second)
		worker.jobId = 0
		wg.Done()
	}
}

func Steamline(workerCount int, jobCount int) {

	var wg sync.WaitGroup
	jobs := make(chan int, jobCount)
	workerList := GetWorkers(workerCount)

	fmt.Println(len(workerList), "WORKERs,", jobCount, "JOBs")

	fmt.Println("START")

	// Monitor Function
	go func() {
		for i := range workerList {
			fmt.Printf("|%s\t", workerList[i].name)
		}
		fmt.Println("|")

		ticker := time.NewTicker(500 * time.Millisecond)
		for range ticker.C {

			for i := range workerList {
				worker := workerList[i]
				jobName := "-"
				if worker.jobId != 0 {
					jobName = strconv.Itoa(worker.jobId)
				}
				fmt.Printf("| %s\t", jobName)
			}
			fmt.Println("|")
		}
	}()

	for i := 0; i < jobCount; i++ {
		jobs <- i+1
		wg.Add(1)
	}

	for i := range workerList {
		go _worker(&workerList[i], jobs, &wg)
	}

	// waiting
	wg.Wait()
	fmt.Println("FINISH")
}
