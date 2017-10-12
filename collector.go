package main

import (
	"github.com/iwilltry42/edgecast"
	"github.com/prometheus/client_golang/prometheus"
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

	// Bandwidth Data for http_small platform
	bw, _ := col.ec.Bandwidth(edgecast.MediaTypeSmall)
	bwBps := bw.Bps
	bwPlatform := platforms[bw.Platform]

	// Bandwidth Data for http_small platform
	bwl, _ := col.ec.Bandwidth(edgecast.MediaTypeLarge)
	bwlBps := bwl.Bps
	bwlPlatform := platforms[bw.Platform]

	ch <- prometheus.MustNewConstMetric(bandwidth, prometheus.GaugeValue, bwBps, []string{bwPlatform}...)
	ch <- prometheus.MustNewConstMetric(bandwidth, prometheus.GaugeValue, bwlBps, []string{bwlPlatform}...)
	ch <- prometheus.MustNewConstMetric(cachestatus, prometheus.GaugeValue, 2, []string{"http_large"}...)
	ch <- prometheus.MustNewConstMetric(connections, prometheus.GaugeValue, 3, []string{"http_large"}...)
	ch <- prometheus.MustNewConstMetric(statuscodes, prometheus.GaugeValue, 4, []string{"http_large"}...)
}
