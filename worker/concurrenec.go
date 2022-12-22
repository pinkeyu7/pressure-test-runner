package worker

import (
	"fmt"
	"pressure-test-runner/api_request"
	"pressure-test-runner/config"
	"pressure-test-runner/file"
	"sync"
	"time"
)

func Concurrence(requests []RequestDto) {
	start := time.Now()
	requestLength := len(requests)

	workerChan := make(chan Worker)
	count := 0
	fileText := ""

	var wg sync.WaitGroup

	for i, request := range requests {
		i := i
		description := request.Description

		wg.Add(1)
		go func() {
			defer wg.Done()
			timeDuration, err := api_request.Get(config.GetApiUrl())
			worker := Worker{
				I:           i,
				Description: description,
				Duration:    timeDuration,
				Err:         err,
			}

			workerChan <- worker
		}()
	}

Loop:
	for worker := range workerChan {
		fmt.Println("i:", worker.I, ", text:", worker.Description, ", time duration:", worker.Duration, "ms")
		fileText += fmt.Sprintf("%d,%s,%d\n", worker.I, worker.Description, worker.Duration)
		count++
		if count == len(requests) {
			break Loop
		}
	}

	wg.Wait()

	end := time.Now()
	timeDiff := end.Sub(start).Milliseconds()

	fileText += fmt.Sprintf("total api_request,%d\n", requestLength)
	fileText += fmt.Sprintf("time duration ,%dms\n", timeDiff)

	// Write file
	logFile := fmt.Sprintf("%d_concurrence_%d_%s.csv", requestLength, timeDiff, time.Now().Format("20060102150405"))
	title := "number,description,duration time,error message\n"
	err := file.Logs(logFile, title, fileText)
	if err != nil {
		panic(err)
	}

	fmt.Println("All Done.", timeDiff, "ms")
}
