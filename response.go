package webque

// StatusMessage struct
type StatusMessage struct {
	Message   string `json:"message"`
	RequestID int    `json:"requestId,omitempty"`
}

// MessageResponse struct
type MessageResponse struct {
	Data StatusMessage `json:"data"`
}

// Account struct
type Account struct {
	ID   int    `json:"accountID"`
	Name string `json:"name"`
}

// LoadRequestResponse struct
type LoadRequestResponse struct {
	Data []LoadRequestModel `json:"data"`
}
