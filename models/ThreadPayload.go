package models

type ThreadPayload struct {
	Messages []Message `json:"messages"`
}
