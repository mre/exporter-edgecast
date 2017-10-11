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
