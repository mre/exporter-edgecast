package main

import (
	"github.com/iwilltry42/edgecast"
	"github.com/prometheus/client_golang/prometheus"
	"sync"
)

// interface to be used for logging and instrumenting middleware
type EdgecastInterface interface {
	Bandwidth(int) (*edgecast.BandwidthData, error)
	Connections(int) (*edgecast.ConnectionData, error)
	CacheStatus(int) (*edgecast.CacheStatusData, error)
	StatusCodes(int) (*edgecast.StatusCodeData, error)
}

type edgecastCollector struct {
	ec EdgecastInterface
}

const (
	NAMESPACE = "Edgecast"
)

var (
	// media-types/platforms
	platforms = map[int]string{
		2:  "flash",
		3:  "http_large",
		8:  "http_small",
		14: "adn",
	}

	// Prepared Description of all fetchable metrics
	bandwidth = prometheus.NewDesc(
		prometheus.BuildFQName(NAMESPACE, "metrics", "bandwidth"), "Bandwidth (Mbps).", []string{"platform"}, nil,
	)
	cachestatus = prometheus.NewDesc(
		prometheus.BuildFQName(NAMESPACE, "metrics", "cachestatus"), "Connections per Cachestatus.", []string{"platform", "CacheStatus"}, nil,
	)
	connections = prometheus.NewDesc(
		prometheus.BuildFQName(NAMESPACE, "metrics", "connections"), "Number of Connections.", []string{"platform"}, nil,
	)
	statuscodes = prometheus.NewDesc(
		prometheus.BuildFQName(NAMESPACE, "metrics", "statuscodes"), "Connections per Statuscode.", []string{"platform", "StatusCode"}, nil,
	)
)

func NewEdgecastCollector(edgecast2 *EdgecastInterface) *edgecastCollector {
	return &edgecastCollector{ec: *edgecast2}
}

/*
 * describes all exported metrics
 * implements function of interface prometheus.Collector
 */
func (col edgecastCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- bandwidth
	ch <- cachestatus
	ch <- connections
	ch <- statuscodes
}

/*
 * fetches metrics and exposes them in Prometheus format
 * implements function of interface prometheus.Collector
 */
func (col edgecastCollector) Collect(ch chan<- prometheus.Metric) {
	var collectWaitGroup sync.WaitGroup
	for p := range platforms { // for each possible platform concurrently
		collectWaitGroup.Add(1)
		go col.metrics(ch, &collectWaitGroup, p) // fetch all possible metrics concurrently
	}
	collectWaitGroup.Wait()
}

func (col edgecastCollector) metrics(ch chan<- prometheus.Metric, collectWaitgroup *sync.WaitGroup, platform int) {
	var metricsWaitGroup sync.WaitGroup
	metricsWaitGroup.Add(4) // 4 goroutines per platform for the 4 possible metric types
	go col.bandwidth(ch, &metricsWaitGroup, platform)
	go col.connections(ch, &metricsWaitGroup, platform)
	go col.cachestatus(ch, &metricsWaitGroup, platform)
	go col.statuscodes(ch, &metricsWaitGroup, platform)
	metricsWaitGroup.Wait() // wait for metric-fetching to finish
	collectWaitgroup.Done() // DONE fetching and exposing metrics for this platform
}

func (col edgecastCollector) bandwidth(ch chan<- prometheus.Metric, metricsWaitGroup *sync.WaitGroup, platform int) {
	defer metricsWaitGroup.Done()

	bw, err := col.ec.Bandwidth(platform)
	if err == nil {
		bwBps := bw.Bps
		bwPlatform := platforms[bw.Platform]
		ch <- prometheus.MustNewConstMetric(bandwidth, prometheus.GaugeValue, bwBps, []string{bwPlatform}...)
	}
}

func (col edgecastCollector) connections(ch chan<- prometheus.Metric, metricsWaitGroup *sync.WaitGroup, platform int) {
	defer metricsWaitGroup.Done()

	con, err := col.ec.Connections(platform)
	if err == nil {
		conCon := con.Connections
		conPlatform := platforms[con.Platform]
		ch <- prometheus.MustNewConstMetric(connections, prometheus.GaugeValue, conCon, []string{conPlatform}...)
	}
}

func (col edgecastCollector) cachestatus(ch chan<- prometheus.Metric, metricsWaitGroup *sync.WaitGroup, platform int) {
	defer metricsWaitGroup.Done()

	cs, err := col.ec.CacheStatus(platform)
	if err == nil {
		csList := *cs
		var val float64
		var labelVals []string
		for c := range csList {
			val = float64(csList[c].Connections)
			labelVals = []string{platforms[platform], csList[c].CacheStatus}
			ch <- prometheus.MustNewConstMetric(cachestatus, prometheus.GaugeValue, val, labelVals...)
		}

	}

}

func (col edgecastCollector) statuscodes(ch chan<- prometheus.Metric, metricsWaitGroup *sync.WaitGroup, platform int) {
	defer metricsWaitGroup.Done()

	sc, err := col.ec.StatusCodes(platform)
	if err == nil {
		scList := *sc
		var val float64
		var labelVals []string
		for s := range scList {
			val = float64(scList[s].Connections)
			labelVals = []string{platforms[platform], scList[s].StatusCode}
			ch <- prometheus.MustNewConstMetric(statuscodes, prometheus.GaugeValue, val, labelVals...)
		}
	}
}
