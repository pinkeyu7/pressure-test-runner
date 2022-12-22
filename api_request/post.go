package api_request

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func Post(apiUrl string) (timeDiff int64, err error) {
	start := time.Now()

	// Handle request dto
	dto := RequestDto{}

	payload, err := json.Marshal(dto)
	if err != nil {
		end := time.Now()
		timeDiff = end.Sub(start).Milliseconds()
		return timeDiff, err
	}

	request, err := http.NewRequest(http.MethodPost, apiUrl, bytes.NewBuffer(payload))
	if err != nil {
		return time.Now().Sub(start).Milliseconds(), err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return time.Now().Sub(start).Milliseconds(), err
	}
	defer response.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		return time.Now().Sub(start).Milliseconds(), err
	}

	if response.StatusCode != http.StatusOK {
		return time.Now().Sub(start).Milliseconds(), errors.New(fmt.Sprintf("api request error, status code:%d", response.StatusCode))
	} else {
		res := ResponseDto{}
		err = json.Unmarshal(buf.Bytes(), &res)
		if err != nil {
			return time.Now().Sub(start).Milliseconds(), err
		}
	}

	return time.Now().Sub(start).Milliseconds(), nil
}
