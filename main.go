package main

import (
	"log"
	"pressure-test-runner/config"
	"pressure-test-runner/file"
	"pressure-test-runner/worker"
	"time"
)

func main() {
	err := file.Clean()
	if err != nil {
		log.Panic(err)
	}

	// Add your api_request data here
	requests := []worker.RequestDto{
		{
			Description: "test for demo",
		},
	}

	for i := 0; i < config.GetTestTimes(); i++ {
		worker.Concurrence(requests)
		time.Sleep(1 * time.Second)
		worker.Queue(requests)
		time.Sleep(1 * time.Second)
	}
	
	file.ReadFile()
}
