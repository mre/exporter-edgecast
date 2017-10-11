package main

import (
	"github.com/go-kit/kit/log"
	ec "github.com/iwilltry42/edgecast"
	"time"
)

type loggingMiddleware struct {
	logger log.Logger
	next   ec.Edgecast
}

/*
 * functions to implement EdgecastInterface
 */

func (mw loggingMiddleware) Bandwidth(platform int) (bandwidthData *ec.BandwidthData, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log( // params: alternating key-value-key-value-...
			"method", "Bandwidth",
			"input", platform,
			"output", bandwidthData,
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
			"input", platform,
			"output", connectionData,
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
			"input", platform,
			"output", cacheStatusData,
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
			"input", platform,
			"output", statusCodeData,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	statusCodeData, err = mw.next.StatusCodes(platform) // hand function call to service
	return
}
