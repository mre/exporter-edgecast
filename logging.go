package main

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/log"
	ec "github.com/iwilltry42/edgecast"
)

type loggingMiddleware struct {
	logger log.Logger
	next   EdgecastInterface
}

/*
 * functions to implement EdgecastInterface
 */

func (mw loggingMiddleware) Bandwidth(platform int) (bandwidthData *ec.BandwidthData, err error) {

	defer func(begin time.Time) {
		_ = mw.logger.Log( // params: alternating key-value-key-value-...
			"method", "Bandwidth",
			"platform", fmt.Sprintf("%d(%s)", platform, Platforms[platform]),
			"output", fmt.Sprintf("%+v", bandwidthData),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	bandwidthData, err = mw.next.Bandwidth(platform) // hand function call to service
	return
}

func (mw loggingMiddleware) Connections(platform int) (connectionData *ec.ConnectionData, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log( // params: alternating key-value-key-value-...
			"method", "Connections",
			"platform", fmt.Sprintf("%d(%s)", platform, Platforms[platform]),
			"output", fmt.Sprintf("%+v", connectionData),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	connectionData, err = mw.next.Connections(platform) // hand function call to service
	return
}

func (mw loggingMiddleware) CacheStatus(platform int) (cacheStatusData *ec.CacheStatusData, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log( // params: alternating key-value-key-value-...
			"method", "CacheStatus",
			"platform", fmt.Sprintf("%d(%s)", platform, Platforms[platform]),
			"output", fmt.Sprintf("%+v", cacheStatusData),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	cacheStatusData, err = mw.next.CacheStatus(platform) // hand function call to service
	return
}

func (mw loggingMiddleware) StatusCodes(platform int) (statusCodeData *ec.StatusCodeData, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log( // params: alternating key-value-key-value-...
			"method", "StatusCodes",
			"platform", fmt.Sprintf("%d(%s)", platform, Platforms[platform]),
			"output", fmt.Sprintf("%+v", statusCodeData),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	statusCodeData, err = mw.next.StatusCodes(platform) // hand function call to service
	return
}
