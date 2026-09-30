package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

var address = os.Getenv("IMAGE_GENERATOR_URL")

func HandleHealth() string {

	response, err := http.Get(address + "health")

	if err != nil {
		return "Cound't access image generator"
	}

	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		return "Error reading data"
	}

	var data map[string]any
	err = json.Unmarshal(responseData, &data)

	if err != nil {
		return "Unable to access data"
	}

	statusVal, ok := data["health"]
	if !ok {
		return "Health status field missing"
	}

	status, ok := statusVal.(string)
	if !ok {
		return "Invalid health status format"
	}

	status = data["health"].(string)
	return status
}
