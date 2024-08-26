package select_chapter

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("Compares speed of servers, returns the url of the faster one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		// Sometimes you need to clean up resources, such as closing a file or
		// closing a server so that it does not continue to listen to a port.
		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		// got := RacerDepricated(slowUrl, fastUrl)

		// Not using ConfigurableRacer because we don't care about timeout requirements in the happy test
		got, err := Racer(slowUrl, fastUrl)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if want != got {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	// Not being kind to my test
	// I've not removed this test suite for my own learning/reference.
	t.Run("returns an error if the server doesn't respond withing 10s", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Second)
		serverB := makeDelayedServer(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		// We would have to actually wait here, which we don't want to do in a unit test.
		// We are being rude to our unit test here.
		// Hence, the improved(kinder) test written below.
		_, err := Racer(serverA.URL, serverB.URL)
		if err == nil {
			t.Errorf("Expected an error but didn't get one")
		}
	})

	// being kind to my test
	t.Run("returns an error if a server does not respond within specified time", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)
		defer server.Close()

		// Using the same url twice really doesn't matter here.
		// I'm being kind by not creating two slow servers, one is enough to test timeout criteria.
		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	// http.HandlerFunc is a type that's used as an adapter
	// that converts a function of a given signature to its type
	// slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	time.Sleep(20 * time.Millisecond)
	// 	w.WriteHeader(http.StatusOK)
	// }))

	// httptest.NewServer returns a test http server that listens on a system-chosen port
	// on the local loopback interface, for end-to-end HTTP tests.
	// It finds an open port to listen on.
	// We close the server when we're done with our test.
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
