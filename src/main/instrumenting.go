package main

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter   // positive only counting value
	requestLatency metrics.Histogram // bucket sampling
	countResult    metrics.Histogram // bucket sampling
	next           EdgecastService
}

func (mw instrumentingMiddleware) GetData(s string) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "getdata", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.GetData(s) // hand request to logged service
	return
}
