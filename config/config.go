package config

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

func GetApiUrl() string {
	return os.Getenv("API_URL")
}

func GetOutputFilePath() string {
	return os.Getenv("OUTPUT_FILE_PATH")
}

func GetTestTimes() int {
	times, err := strconv.Atoi(os.Getenv("TEST_TIMES"))
	if err != nil {
		return 0
	}
	return times
}
