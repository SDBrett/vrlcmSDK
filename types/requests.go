package types

type Request struct{
	RequestID string `json:"requestId"`
	RequestState string `json:"requestState, omitempty"`
	RequestData string `json:"requestData, omitempty"`
	RequestDataType string `json:"requestDataType, omitempty"`
	ExecutionPath string `json:"executionPath, omitempty"`
	RequestStatus string `json:"requestStatus, omitempty"`
	RequestType string `json:"requestType, omitempty"`
	ErrorCause string `json:"errorCause, omitempty"`
	RequestUserID string `json:"requestUserID, omitempty"`
	DocumentVersion string `json:"documentVersion, omitempty"`
	DocumentKind string `json:"documentKind, omitempty"`
	DocumentSelfLink string `json:"documentSelfLink, omitempty"`
	DocumentUpdateTimeMicros string `json:"documentUpdateTimeMicros, omitempty"`
	DocumentUpdateAction string `json:"documentUpdateAction, omitempty"`
	DocumentExpirationTimeMicros string `json:"documentExpirationTimeMicros, omitempty"`
	DocumentAuthPrincipalLink string `json:"documentAuthPrincipalLink, omitempty"`
}

type RequestRetryParameters struct {
	EventId   string `json:"eventId"`
	MessageId string `json:"messageId"`
	Message   string `json:"message"`
}

type VCenterRequest struct {
	RequestID       string                   `json:"id"`
	Type            string                   `json:"type"`
	State           string                   `json:"state"`
	Status          string                   `json:"status"`
	IsRetriable     bool                     `json:"isRetriable"`
	RetryParameters []RequestRetryParameters `json:"retryParameters"`
}