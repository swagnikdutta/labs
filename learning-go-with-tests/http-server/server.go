package main

import (
	"fmt"
	"net/http"
	"strings"
)

type iPlayerStore interface {
	GetPlayerScore(name string) int
}

// The server and its handlers

type PlayerServer struct {
	store iPlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		w.WriteHeader(http.StatusAccepted)
		return
	}

	player := strings.TrimPrefix(r.URL.Path, "/players/")

	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}
