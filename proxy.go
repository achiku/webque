package webque

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/xhandler"
	"github.com/rs/xmux"

	"golang.org/x/net/context"
)

// LoadRequest add load request
func LoadRequest(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	res := MessageResponse{Data: StatusMessage{Message: "load request"}}
	json.NewEncoder(w).Encode(res)
}

// GetLoadRequest get load requests
func GetLoadRequest(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	res := MessageResponse{Data: StatusMessage{Message: "get load request"}}
	json.NewEncoder(w).Encode(res)
}

// TransferRequest add transfer request
func TransferRequest(ctx context.Context, w http.ResponseWriter, r *http.Request) {
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
	context.WithValue(rootCtx, ctxKeyDB, db)

	mux := xmux.New()
	mux.NotFound = xhandler.HandlerFuncC(NotFound)

	api := mux.NewGroup("/api")
	api.POST("/load", xhandler.HandlerFuncC(LoadRequest))
	api.POST("/transfer", xhandler.HandlerFuncC(TransferRequest))
	api.GET("/load", xhandler.HandlerFuncC(GetLoadRequest))
	api.GET("/transfer", xhandler.HandlerFuncC(GetTransferRequest))
	api.GET("/deposit", xhandler.HandlerFuncC(GetCurrentDeposit))

	if err := http.ListenAndServe(":8899", c.HandlerCtx(rootCtx, mux)); err != nil {
		log.Fatal(err)
	}
}
