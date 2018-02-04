package main

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"github.com/castaneai/mf"
	"os"
	"net/http"
	"context"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

func init() {
	r := gin.Default()
	r.GET("/total", getTotalAsset)
	http.Handle("/", r)
}

func handleError(ctx context.Context, err error) {
	log.Errorf(ctx, "%+v", err)
}

func getTotalAsset(c *gin.Context) {
	ctx := appengine.NewContext(c.Request)

	hc := urlfetch.Client(ctx)
	opts := &mf.ClientOption{Host: "https://moneyforward.com", SessionID: os.Getenv("MF_SESSION")}
	mfc, err := mf.NewClient(hc, opts)
	if err != nil {
		handleError(ctx, err)
		return
	}

	totalAsset, err := mfc.GetTotalAsset(ctx)
	if err != nil {
		handleError(ctx, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"amount": totalAsset.Amount,
	})
}
