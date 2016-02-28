package webque

import "time"

// LoadRequestModel load request struct
type LoadRequestModel struct {
	ID        int32     `json:"id"`
	AccountID int32     `json:"accountID"`
	Amount    int32     `json:"amount"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
