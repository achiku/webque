package webque

import "github.com/jackc/pgx"

// LoadRequestService create load request
func LoadRequestService(tx *pgx.Tx, req LoadRequestRequest) (int, error) {
	return 1, nil
}
