package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

type Job struct {
	id       int
	randomno int
}

type Result struct {
	job         Job
	sumofdigits int
}

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		fmt.Println(job)
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	fmt.Printf("worker %v is done.\n", wg)
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		fmt.Println(i)
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func main() {
	// Sends to a buffered channel are blocked only when the buffer is full.
	// Similarly receives from a buffered channel are blocked only when the buffer is empty.
	/* ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	fmt.Println(<-ch)
	fmt.Println(<-ch) */

	/* ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)
	} */

	/* ch := make(chan string, 3)
	ch <- "naveen"
	ch <- "paul"
	fmt.Println("capacity is", cap(ch))
	fmt.Println("length is", len(ch))
	fmt.Println("read value", <-ch)
	fmt.Println("new length is", len(ch)) */

	/* no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing") */

	startTime := time.Now()
	noOfJobs := 20
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 5
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
