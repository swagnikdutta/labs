package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func nameOfWeekday(d int) string {
	var day string
	switch d {
	case 0:
		day = "monday"
	case 1:
		day = "tuesday"
	case 2:
		day = "wednesday"
	case 3:
		day = "thursday"
	case 4:
		day = "friday"
	case 5:
		day = "saturday"
	case 6:
		day = "sunday"
	default:
		return "invalid"
	}
	return day
}

func reverseInPlace(integers []int) {
	l, r := 0, len(integers)-1
	for l < r {
		temp := integers[l]
		integers[l] = integers[r]
		integers[r] = temp
		l, r = l+1, r-1
	}
}

func simulateClosingTask() {
	fmt.Println("all systems shutdown successfully")
}

func demoDefer() {
	// stuff deferred will be executed last, before this function returns.
	defer simulateClosingTask()
	fmt.Println()
	fmt.Println("shutting down all systems in 1 second")
	time.Sleep(1 * time.Second)
}

func pingpong() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan string)

	var pingCount int32 = 0

	// producer side
	go func() {
		defer wg.Done()
		for {
			if atomic.AddInt32(&pingCount, 1) > 5 {
				close(ch) // close the channel from the producer side
				return
			}
			ch <- "ping"
			time.Sleep(1 * time.Second)
		}
	}()

	// consumer side
	go func() {
		defer wg.Done()
		for msg := range ch {
			fmt.Println(msg)
		}
	}()

	wg.Wait()
}

// sending to channel. So send-only channel
func sendOnly(ch chan<- string) {
	ch <- "Hello"
	ch <- "from"
	ch <- "sender"
	close(ch)
}

// receiving from channel, so receive-only channel
func receiveOnly(ch <-chan string) {
	for msg := range ch {
		fmt.Println("Received: ", msg)
	}
}

func workerPoolBufferedChannel() {
	// The most important thing is to ensure, all of your jobs are done.
	// One way of ensuring that use to add jobCount to a WaitGroup
	numJobs := 20
	numWorkers := 3

	// here we use a jobs channel with size equal to the number of jobs.
	// The buffer size can be anything really. If it's smaller than the
	// number of jobs, the producer code will block for some time until
	// there's space in the buffer again.
	// CONCLUSION: Buffer size does not affect correctness — only performance.
	// If it's smaller than numJobs, the producer may block temporarily when the buffer is full.
	jobsChannel := make(chan int, numJobs)

	wg := sync.WaitGroup{}

	// get the workers ready
	for i := 1; i <= numWorkers; i++ {
		// creating workers
		go func(workerId int) {
			for jobId := range jobsChannel {
				fmt.Printf("Worker %d is starting job %d\n", workerId, jobId)
				time.Sleep(1 * time.Second)
				fmt.Printf("Worker %d has finished job %d\n", workerId, jobId)
				wg.Done()
			}
		}(i)
	}

	// add the total number of jobs in the WaitGroup, and start producing jobs
	wg.Add(numJobs)
	for i := 1; i <= numJobs; i++ {
		jobsChannel <- i
	}

	// Closing the channel, as otherwise the goroutines will wait/listen forever on the jobs channel.
	// The goroutines will leak silently.
	//
	// Silently because the program wouldn't crash. The WaitGroup's counter reaches 0 because it's
	// only decremented when each job is done — not when the goroutine exits.
	//
	// But if the WaitGroup count were tracking the *lifetime of the goroutines themselves*
	// (instead of just the jobs), then not closing the channel would have caused a hang —
	// because the goroutines would never exit, and wg.Wait() would block forever.
	//
	// Either way, once there are no jobs left to assign, it’s the right thing to do to let the
	// workers know they’re done for the day. Close the channel — that’s like putting the phone
	// down and letting the workers go home.
	close(jobsChannel)

	wg.Wait()
}

func workerPoolUnbufferedChannel() {
	numJobs := 20
	numWorkers := 3

	// Since we use an unbuffered channel this time, we would make the producer — who's trying
	// to push jobs into this channel, wait a lot more than in the above approach. We are putting
	// some backpressure on the producer.
	jobsChannel := make(chan int)

	wg := sync.WaitGroup{}
	// In this approach, WaitGroup is tied to number of worker goroutines.
	wg.Add(numWorkers)

	// prepare the workers
	for i := 1; i <= numWorkers; i++ {
		go func(workerId int) {
			defer wg.Done()

			for jobId := range jobsChannel {
				fmt.Printf("Worker %d is starting job %d\n", workerId, jobId)
				time.Sleep(1 * time.Second)
				fmt.Printf("Worker %d has finished job %d\n", workerId, jobId)
			}
		}(i)
	}

	// producer
	for i := 1; i <= numJobs; i++ {
		jobsChannel <- i
	}

	// Forget to close this channel, and the program would crash — not silently.
	close(jobsChannel)

	wg.Wait()
}

func main() {
	// function to reverse slice of ints in-place
	integers := []int{10, 20, 30, 40, 50}
	reverseInPlace(integers)
	fmt.Println(integers)

	for i := 0; i < 7; i++ {
		fmt.Print(nameOfWeekday(i), " ")
	}

	demoDefer()
	// pingpong()

	ch := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(2)

	// This wrapper goroutine is a great way to use WaitGroup,
	// and yet not pass it to the actual function,
	// keeping its arguments list minimal.
	go func() {
		defer wg.Done()
		sendOnly(ch)
	}()

	go func() {
		defer wg.Done()
		receiveOnly(ch)
	}()

	wg.Wait()

	// doWorkerPoolAgain()
	workerPoolBufferedChannel()
	workerPoolUnbufferedChannel()
}
