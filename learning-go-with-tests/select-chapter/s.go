package select_chapter

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func measureResponseTime(url string) time.Duration {
	startTime := time.Now()
	http.Get(url)
	return time.Since(startTime)
}

func RacerDepricated(a, b string) string {
	aDuration := measureResponseTime(a)
	bDuration := measureResponseTime(b)

	if aDuration < bDuration {
		return a
	}
	return b
}

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	// A channel just for signalling purposes. We don't care
	// what type we send to the channel.
	//
	// Why not use bool instead?
	//
	// An empty struct{} is the smallest possible datatype in Go. It requires no memeory allocation
	// while a boolean does(1 byte).
	//
	// With `chan bool`, even if we simply closed the channel without sending any value through it,
	// the type of the channel would still allocate memory for each boolean value it can hold.
	//
	// While `chan struct{}` requires no additional memory for the elements sent through the channel
	// because, the element (`struct{}`) itself occupies no memory.
	//
	// Hence, the latter becomes the perfect choice when the value you send throught the channel is
	// irrelvant and the act of sending or closing the channel is purely for signalling purposes.

	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
