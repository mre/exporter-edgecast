package main

import (
	"github.com/go-kit/kit/log"
	ec "github.com/iwilltry42/edgecast"
	"time"
)

type loggingMiddleware struct {
	logger log.Logger
	next   EdgecastService
}

// logger for GetData function of edgecastService
func (mw loggingMiddleware) GetData(s string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log( // params: alternating key-value-key-value-...
			"method", "getdata",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.GetData(s) // hand function call to service
	return
}

func (mw loggingMiddleware) GetBandwidth(platform int) (bandwidthData *ec.BandwidthData, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log( // params: alternating key-value-key-value-...
			"method", "getbandwidth",
			"input", platform,
			"output", bandwidthData,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	bandwidthData, err = mw.next.GetBandwidth(platform) // hand function call to service
	return
}
