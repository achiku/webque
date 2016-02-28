package webque

// GetLoadRequestRequest load request
type GetLoadRequestRequest struct {
	AccountID int `json:"accountId"`
}

// LoadRequestRequest load request struct
type LoadRequestRequest struct {
	AccountID int `json:"accountId"`
	Amount    int `json:"amount"`
}

// TransferRequestRequest transfer request struct
type TransferRequestRequest struct {
	FromAccountID int `json:"fromAccountId"`
	ToAccountID   int `json:"toAccountId"`
	Amount        int `json:"amount"`
}
