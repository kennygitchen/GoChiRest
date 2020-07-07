package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"time"
	"log"
	"math/rand"
	"fmt"
)

func main() {
    start := time.Now()
    random := rand.New(rand.NewSource(99))

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/rand", func(w http.ResponseWriter, r *http.Request) {
        number := random.Float64()
		w.Write([]byte(fmt.Sprintf("%f",number)))
	})

    done := make(chan bool)
	go http.ListenAndServe(":8080", r)
	timeTrack(start, "Service Start up on port 8080")
	<-done
}

func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    log.Printf("%s taken %d(ms)", name, elapsed/1000)
}