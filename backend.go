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

// UpdateCurrentDeposit create/update current deposit
func UpdateCurrentDeposit(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	res := MessageResponse{Data: StatusMessage{Message: "current deposit"}}
	json.NewEncoder(w).Encode(res)
}

// BackendRun runs backend server
func BackendRun() {
	fmt.Println("starting backend service...")

	db, err := NewBackendDB("postgresql://localhost/webque_backend")
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
	api.POST("/deposit", xhandler.HandlerFuncC(UpdateCurrentDeposit))

	if err := http.ListenAndServe(":8889", c.HandlerCtx(rootCtx, mux)); err != nil {
		log.Fatal(err)
	}
}
