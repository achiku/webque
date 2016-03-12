package webque

import (
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
	log.Println("starting proxy service...")
	db, err := NewProxyDB("postgresql://localhost/webque_proxy")
	if err != nil {
		log.Fatal(err)
	}
	qc, err := NewQueClient("postgresql://localhost/webque_proxy")
	if err != nil {
		log.Fatal(err)
	}
	wm := que.WorkMap{
		"HelloJob":         HelloJob,
		"UpdateDepositJob": UpdateDepositJob,
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
	api.POST("/load/request/:accountID", xhandler.HandlerFuncC(CreateLoadRequest))
	api.PUT("/load/request/:requestID", xhandler.HandlerFuncC(CompleteLoadRequest))
	api.POST("/transfer/request/:accountID", xhandler.HandlerFuncC(CreateTransferRequest))
	api.GET("/load/request/:accountID", xhandler.HandlerFuncC(GetLoadRequest))
	api.GET("/transfer/request/:accountID", xhandler.HandlerFuncC(GetTransferRequest))
	api.GET("/deposit/:accountID", xhandler.HandlerFuncC(GetCurrentDeposit))

	if err := http.ListenAndServe(":8899", c.HandlerCtx(rootCtx, mux)); err != nil {
		log.Fatal(err)
	}
}
