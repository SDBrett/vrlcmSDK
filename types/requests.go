package types

type Request struct {
	ID              string                   `json:"id"`
	Type            string                   `json:"type"`
	State           string                   `json:"state"`
	Status          string                   `json:"status"`
	IsRetriable     bool                     `json:"isRetriable"`
	RetryParameters []RequestRetryParameters `json:"retryParameters"`
}

type RequestRetryParameters struct {
	EventId   string `json:"eventId"`
	MessageId string `json:"messageId"`
	Message   string `json:"message"`
}
