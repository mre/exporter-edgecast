package main

import (
	prometheus "github.com/prometheus/client_golang/prometheus"
)

type edgecastCollector struct{}

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

func (ec edgecastCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- bandwidth
	ch <- cachestatus
	ch <- connections
	ch <- statuscodes
}

func (ec edgecastCollector) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(bandwidth, prometheus.GaugeValue, 2, []string{"platform"}...)
}
