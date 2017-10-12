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

	// possible variableLabels for metrics exposed to prometheus
	variableLabels = []string{"platform"}

	// Prepared Description of all fetchable metrics
	bandwidth = prometheus.NewDesc(
		prometheus.BuildFQName(NAMESPACE, "metrics", "bandwidth"), "Bandwidth (Mbps).", variableLabels, nil,
	)
	cachestatus = prometheus.NewDesc(
		prometheus.BuildFQName(NAMESPACE, "metrics", "cachestatus"), "Connections per Cachestatus.", variableLabels, nil,
	)
	connections = prometheus.NewDesc(
		prometheus.BuildFQName(NAMESPACE, "metrics", "connections"), "Number of Connections.", variableLabels, nil,
	)
	statuscodes = prometheus.NewDesc(
		prometheus.BuildFQName(NAMESPACE, "metrics", "statuscodes"), "Connections per Statuscode.", variableLabels, nil,
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
	bw, _ := col.ec.Bandwidth(platform)
	bwBps := bw.Bps
	bwPlatform := platforms[bw.Platform]
	ch <- prometheus.MustNewConstMetric(bandwidth, prometheus.GaugeValue, bwBps, []string{bwPlatform}...)
	metricsWaitGroup.Done()
}

func (col edgecastCollector) connections(ch chan<- prometheus.Metric, metricsWaitGroup *sync.WaitGroup, platform int) {
	con, _ := col.ec.Connections(platform)
	conCon := con.Connections
	conPlatform := platforms[con.Platform]
	ch <- prometheus.MustNewConstMetric(connections, prometheus.GaugeValue, conCon, []string{conPlatform}...)
	metricsWaitGroup.Done()
}

func (col edgecastCollector) cachestatus(ch chan<- prometheus.Metric, metricsWaitGroup *sync.WaitGroup, platform int) {
	// TODO: do something
	metricsWaitGroup.Done()
}

func (col edgecastCollector) statuscodes(ch chan<- prometheus.Metric, metricsWaitGroup *sync.WaitGroup, platform int) {
	// TODO: do something
	metricsWaitGroup.Done()
}
