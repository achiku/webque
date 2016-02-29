package webque

import (
	"fmt"
	"log"
	"net/http"
	"time"

	que "github.com/bgentry/que-go"
	"github.com/rs/xhandler"
	"github.com/rs/xmux"

	"golang.org/x/net/context"
)

// ProxyRun runs proxy service
func ProxyRun() {
	fmt.Println("starting proxy service...")

	db, err := NewDB("postgresql://localhost/webque_proxy")
	if err != nil {
		log.Fatal(err)
	}
	qc, err := NewQueClient("postgresql://localhost/webque_proxy")
	if err != nil {
		log.Fatal(err)
	}
	wm := que.WorkMap{
		"HelloJob": HelloJob,
	}
	log.Println("create worker pool")
	workers := que.NewWorkerPool(qc, wm, 2)
	workers.Interval = 2 * time.Second
	go workers.Start()

	c := xhandler.Chain{}
	c.Use(recoverMiddleware)
	c.Use(loggingMiddleware)
	c.Use(jsonResponseMiddleware)
	rootCtx := context.Background()
	rootCtx = context.WithValue(rootCtx, ctxKeyDB, db)
	rootCtx = context.WithValue(rootCtx, ctxKeyQueClient, qc)

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
