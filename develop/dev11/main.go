package main

import (
	"log"
	"sync"

	"dev11/handler"
	"dev11/middleware"
	"dev11/repos"
	"dev11/server"
)

func main() {
	repo := repos.NewStore(&sync.Mutex{}, make(map[int]repos.Event))

	srv := server.New()

	handl := handler.New(repo).InitRoutes()

	mw := middleware.RequestLogging(handl)

	log.Fatal(srv.Run("8080", mw))
}
