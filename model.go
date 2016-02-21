package webque

// StatusMessage struct
type StatusMessage struct {
	Message string `json:"message"`
}

// MessageResponse struct
type MessageResponse struct {
	Data StatusMessage `json:"data"`
}
