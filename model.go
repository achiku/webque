package webque

import "time"

// LoadRequestModel load request struct
type LoadRequestModel struct {
	ID        int
	AccountID int
	Completed bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
