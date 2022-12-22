package worker

import (
	"fmt"
	"pressure-test-runner/api_request"
	"pressure-test-runner/config"
	"pressure-test-runner/file"
	"time"
)

func Queue(requests []RequestDto) {
	fileText := ""
	start := time.Now()
	requestLength := len(requests)

	for i, request := range requests {
		timeDuration, err := api_request.Get(config.GetApiUrl())
		description := request.Description
		worker := Worker{
			I:           i,
			Description: description,
			Duration:    timeDuration,
			Err:         err,
		}

		fmt.Println("i:", worker.I, ", text:", worker.Description, ", time duration:", worker.Duration, "ms")
		fileText += fmt.Sprintf("%d,%s,%d\n", worker.I, worker.Description, worker.Duration)

		time.Sleep(100 * time.Millisecond)
	}

	end := time.Now()
	timeDiff := end.Sub(start).Milliseconds()

	fileText += fmt.Sprintf("total api_request,%d\n", requestLength)
	fileText += fmt.Sprintf("time duration ,%dms\n", timeDiff)

	// Write file
	logFile := fmt.Sprintf("%d_queue_%d_%s.csv", requestLength, timeDiff, time.Now().Format("20060102150405"))
	title := "number,duration time,error message\n"
	err := file.Logs(logFile, title, fileText)
	if err != nil {
		panic(err)
	}

	fmt.Println("All Done.", timeDiff, "ms")
}
