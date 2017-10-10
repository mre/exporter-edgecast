package main

import (
	prometheus "github.com/prometheus/client_golang/prometheus"
)

type collector struct{}

const (
	NAMESPACE = "Edgecast"
)

var (
	variableLabels = []string{"platform"}
	bandwidth      = prometheus.NewDesc(
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

/*
 * describes all exported metrics
 * implements function of interface prometheus.Collector
 */
func (ec collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- bandwidth
	ch <- cachestatus
	ch <- connections
	ch <- statuscodes
}

/*
 * fetches metrics and exposes them in Prometheus format
 * implements function of interface prometheus.Collector
 */
func (ec collector) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(bandwidth, prometheus.GaugeValue, 2, []string{"platform"}...)
}
