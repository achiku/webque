package webque

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jackc/pgx"
	"github.com/rs/xhandler"
	"github.com/rs/xmux"

	"golang.org/x/net/context"
)

// CreateLoadRequest add load request
func CreateLoadRequest(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	res := MessageResponse{Data: StatusMessage{Message: "request accepted", RequestID: 11}}
	json.NewEncoder(w).Encode(res)
}

// GetLoadRequest get load requests
func GetLoadRequest(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	accountID, _ := strconv.Atoi(xmux.Param(ctx, "accountID"))
	db := ctx.Value(ctxKeyDB).(*pgx.ConnPool)
	tx, _ := db.Begin()
	req := GetLoadRequestRequest{
		AccountID: accountID,
	}
	loadRequests, err := GetLoadRequestService(tx, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := MessageResponse{Data: StatusMessage{Message: "failed"}}
		json.NewEncoder(w).Encode(res)
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

// ProxyRun runs proxy service
func ProxyRun() {
	fmt.Println("starting proxy service...")

	db, err := NewDB("postgresql://localhost/webque_proxy")
	if err != nil {
		log.Fatal(err)
	}

	c := xhandler.Chain{}
	c.Use(recoverMiddleware)
	c.Use(loggingMiddleware)
	c.Use(jsonResponseMiddleware)
	rootCtx := context.Background()
	rootCtx = context.WithValue(rootCtx, ctxKeyDB, db)

	mux := xmux.New()
	mux.NotFound = xhandler.HandlerFuncC(NotFound)

	api := mux.NewGroup("/api")
	api.POST("/load/:accountID", xhandler.HandlerFuncC(CreateLoadRequest))
	api.POST("/transfer/:accountID", xhandler.HandlerFuncC(CreateTransferRequest))
	api.GET("/load/:accountID", xhandler.HandlerFuncC(GetLoadRequest))
	api.GET("/transfer/:accountID", xhandler.HandlerFuncC(GetTransferRequest))
	api.GET("/deposit/:accountID", xhandler.HandlerFuncC(GetCurrentDeposit))

	if err := http.ListenAndServe(":8899", c.HandlerCtx(rootCtx, mux)); err != nil {
		log.Fatal(err)
	}
}
