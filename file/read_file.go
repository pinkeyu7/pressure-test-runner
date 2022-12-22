package file

import (
	"fmt"
	"io/ioutil"
	"pressure-test-runner/config"
	"strconv"
	"strings"
)

func ReadFile() {
	mapping := map[int]map[string][]int{}

	folder := config.GetOutputFilePath()
	csvFiles, err := ioutil.ReadDir(folder)
	if err != nil {
		panic(err)
	}
	for _, csvFile := range csvFiles {
		if csvFile.Name() == ".DS_Store" {
			continue
		}

		parameters := strings.Split(csvFile.Name(), "_")

		length, _ := strconv.Atoi(parameters[0])
		qType := parameters[1]
		timeDiff, _ := strconv.Atoi(parameters[2])

		if mapping[length] == nil {
			mapping[length] = map[string][]int{}
		}
		if mapping[length][qType] == nil {
			mapping[length][qType] = []int{}
		}

		mapping[length][qType] = append(mapping[length][qType], timeDiff)
	}

	fileText := ""
	for lKey, length := range mapping {
		for qKey, qType := range length {
			fileText += fmt.Sprint(lKey, ",", qKey, ",")
			sum := 0
			for _, timeDiff := range qType {
				sum += timeDiff
				fileText += fmt.Sprint(timeDiff, ",")
			}
			fileText += fmt.Sprint(sum, ",", sum/len(qType), "\n")
		}
	}

	numberString := ""
	for i := 0; i < config.GetTestTimes(); i++ {
		numberString += fmt.Sprintf("%d,", i+1)
	}

	logFile := "output.csv"
	title := fmt.Sprintf("requests,type,%ssum,average\n", numberString)
	err = Logs(logFile, title, fileText)
	if err != nil {
		panic(err)
	}
}
