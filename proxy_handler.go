package webque

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bgentry/que-go"
	"github.com/jackc/pgx"
	"github.com/rs/xmux"
	"golang.org/x/net/context"
)

// CreateLoadRequest add load request
func CreateLoadRequest(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	accountID, _ := strconv.Atoi(xmux.Param(ctx, "accountID"))
	db := ctx.Value(ctxKeyDB).(*pgx.ConnPool)
	qc := ctx.Value(ctxKeyQueClient).(*que.Client)
	tx, _ := db.Begin()
	defer tx.Rollback()

	req := LoadRequestRequest{
		AccountID: accountID,
		Amount:    1000,
	}
	if err := CreateLoadRequestService(tx, req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := MessageResponse{Data: StatusMessage{Message: err.Error()}}
		json.NewEncoder(w).Encode(res)
		return
	}
	if err := qc.EnqueueInTx(&que.Job{Type: "HelloJob"}, tx); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := MessageResponse{Data: StatusMessage{Message: err.Error()}}
		json.NewEncoder(w).Encode(res)
		return
	}

	tx.Commit()
	res := MessageResponse{Data: StatusMessage{Message: "request created"}}
	json.NewEncoder(w).Encode(res)
}

// GetLoadRequest get load requests
func GetLoadRequest(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	accountID, _ := strconv.Atoi(xmux.Param(ctx, "accountID"))
	db := ctx.Value(ctxKeyDB).(*pgx.ConnPool)
	tx, _ := db.Begin()
	defer tx.Rollback()

	req := GetLoadRequestRequest{
		AccountID: accountID,
	}
	loadRequests, err := GetLoadRequestService(tx, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := MessageResponse{Data: StatusMessage{Message: "failed"}}
		json.NewEncoder(w).Encode(res)
		return
	}
	res := LoadRequestResponse{Data: loadRequests}
	json.NewEncoder(w).Encode(res)
}

// CreateTransferRequest add transfer request
func CreateTransferRequest(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	res := MessageResponse{Data: StatusMessage{Message: "transfer request"}}
	json.NewEncoder(w).Encode(res)
}

// GetTransferRequest get transfer request
func GetTransferRequest(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	res := MessageResponse{Data: StatusMessage{Message: "get transfer request"}}
	json.NewEncoder(w).Encode(res)
}

// GetCurrentDeposit get current deposit
func GetCurrentDeposit(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	res := MessageResponse{Data: StatusMessage{Message: "get current deposit"}}
	json.NewEncoder(w).Encode(res)
}

// NotFound returns not found message
func NotFound(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	res := MessageResponse{Data: StatusMessage{Message: "not found"}}
	json.NewEncoder(w).Encode(res)
}
