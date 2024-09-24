package main

import (
	"log"

	"github.com/sebastianmendez/trail-browser/internal/app"
)

type App struct {
	signalClose chan error
}

func main() {
	appVal := &App{
		signalClose: make(chan error, 1),
	}

	go app.Start("/", ":80").ListenAndServe(appVal.signalClose)

	err := <-appVal.signalClose
	if err != nil {
		log.Fatalf("HTTP server gone: %s", err)
	}
}
