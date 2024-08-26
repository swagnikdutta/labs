package main

import (
	"os"
	"time"
)

type RealSleeper struct{}

func (r *RealSleeper) Sleep() {
	// wrapping the original time.Sleep call within
	time.Sleep(1 * time.Second)
}

// time for a configurable sleeper

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(duration time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration) // this sleep thing could be a real thing or a fake thing depending on who's invoking it - main or some test
}

func main() {
	// realSleeper := &RealSleeper{}
	// Countdown(os.Stdout, realSleeper)

	veryRealSleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep} // not attaching any fakeness here
	Countdown(os.Stdout, veryRealSleeper)
}
