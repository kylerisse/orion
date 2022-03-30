package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gitlab.com/nebulaworks/eng/bazaar/prj/term-apply/pkg/server"
)

const host = "localhost"
const port = 23234
const uploadDir = "./uploads"
const dataFile = "applicants.csv"

func main() {

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	s, err := server.NewServer(host, uploadDir, dataFile, port)
	if err != nil {
		log.Printf("Cannot create server %v", err)
	}
	s.Start()

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	s.Stop(ctx)

}
