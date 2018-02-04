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
	"google.golang.org/appengine/memcache"
	"fmt"
	"time"
	"strconv"
	"github.com/gin-contrib/cors"
)

const (
	getTotalAssetCacheKey = "totalAsset"
	cacheExpireSeconds = 86400
)

func init() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/total", getTotalAsset)
	http.Handle("/", r)
}

func handleError(ctx context.Context, err error) {
	log.Errorf(ctx, "%+v", err)
}

func getTotalAsset(c *gin.Context) {
	ctx := appengine.NewContext(c.Request)

	item, err := memcache.Get(ctx, getTotalAssetCacheKey)
	if err != nil && err != memcache.ErrCacheMiss {
		handleError(ctx, err)
		return
	}

	if item == nil {
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

		item = &memcache.Item{
			Key: getTotalAssetCacheKey,
			Value: []byte(fmt.Sprintf("%d", totalAsset.Amount)),
			Expiration: cacheExpireSeconds * time.Second,
		}

		if err := memcache.Set(ctx, item); err != nil {
			handleError(ctx, err)
			return
		}
	}

	totalAmount, err := strconv.Atoi(string(item.Value))
	if err != nil {
		handleError(ctx, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"amount": totalAmount,
	})
}
