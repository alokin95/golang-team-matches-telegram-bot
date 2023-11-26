package assistant

import (
	"encoding/json"
	"io"
	"kadigramo/client"
	"kadigramo/models"
	"log"
	"net/http"
)

const createThreadEndpoint = "https://api.openai.com/v1/threads"

func CreateThread(httpClient client.HttpClient, messages []models.Message) (models.Thread, error) {
	threadPayload := models.ThreadPayload{Messages: messages}
	payloadBytes, err := json.Marshal(threadPayload)
	if err != nil {
		log.Printf("Error marshaling payload: %v\n", err)
		return models.Thread{}, err
	}

	resp, err := httpClient.SendPost(createThreadEndpoint, payloadBytes)
	if err != nil {
		log.Printf("Error making POST request: %v\n", err)
		return models.Thread{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return handleNonOkResponse(resp)
	}

	return decodeThread(resp)
}

func handleNonOkResponse(resp *http.Response) (models.Thread, error) {
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v\n", err)
		return models.Thread{}, err
	}
	log.Printf("Non-OK HTTP status: %d\nResponse Body: %s\n", resp.StatusCode, string(bodyBytes))
	return models.Thread{}, err
}

func decodeThread(resp *http.Response) (models.Thread, error) {
	var thread models.Thread
	err := json.NewDecoder(resp.Body).Decode(&thread)
	if err != nil {
		log.Printf("Error decoding response to thread: %v\n", err)
		return models.Thread{}, err
	}
	return thread, nil
}
