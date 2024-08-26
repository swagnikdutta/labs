package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

// Spies are kind of a mock that record how a dependency is used
// Here my dependency is the sleep call, I can record how many times
// it was called

// Whatever it is you're dependent on, if you have to nudge/modify the
// behaviour of this dependency, create an interface of it.
// Eg, if you have to call write, you need to record this call,
// if you have to call sleep, you need to record that call as well.
// Here, recording = nudging
type SpyWriterSleeper struct {
	CallsSequence []string
}

func (s *SpyWriterSleeper) Sleep() {
	s.CallsSequence = append(s.CallsSequence, "sleep")
}

func (s *SpyWriterSleeper) Write(p []byte) (n int, err error) {
	s.CallsSequence = append(s.CallsSequence, "write")
	return
}

func TestCountdown(t *testing.T) {
	t.Run("test order of operations", func(t *testing.T) {
		// In this test we are only concerned about the order of operations.
		// We won't be able to check contents because, even we have our own implementation
		// of writer, we would still need a writer type argumenr - buffer or os.stdout.
		// But the func `Write(p []byte) (n int, err error)` does not take any such argument.
		//
		// Let me try again.
		// In buf.Write() you actually write to the buffer (size of b is increases) and from that
		// we can obtain a string.
		// But in the Write() implementation of SpyWriterSleeper, we only append the operations
		// to the CallSequence slice `[write,sleep,write,sleep]` and thus there is no way
		// to get a string `3\n2\n1\nGo!` from it.

		spyWriterSleeper := &SpyWriterSleeper{}
		Countdown(spyWriterSleeper, spyWriterSleeper)

		want := []string{
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
		}
		got := spyWriterSleeper.CallsSequence

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Call sequence mismatch, want %v, got %v", want, got)
		}
	})
	t.Run("test content", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spyWriterSleeper := &SpyWriterSleeper{}

		Countdown(buffer, spyWriterSleeper)
		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

// defining the fake sleeper
// just so that we can plug the spy sleeper. We defined a type because maybe having a gloabl spy function is
// bad in go.
type FakeSleeper struct {
	timeSlept time.Duration
}

func (s *FakeSleeper) Sleep(d time.Duration) {
	s.timeSlept = d
}

// end of fakeness

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	fakeSleeper := &FakeSleeper{}

	sleeper := ConfigurableSleeper{sleepTime, fakeSleeper.Sleep} // adding a pinch of fakeness
	sleeper.Sleep()

	if fakeSleeper.timeSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v\n", sleepTime, fakeSleeper.timeSlept)
	}
}
