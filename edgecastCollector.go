package main

import (
	prometheus "github.com/prometheus/client_golang/prometheus"
)

type edgecastCollector struct{}

var (
	Namespace = "Edgecast"
	testDesc  = prometheus.NewDesc(
		prometheus.BuildFQName(Namespace, "api_metrics", "bandwidth"), "BandwidthData", []string{"platform"}, nil,
	)
)

func (ec edgecastCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- testDesc
}

func (ec edgecastCollector) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(testDesc, prometheus.GaugeValue, 2, []string{"platform"}...)
}
