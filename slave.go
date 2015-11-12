package main

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func newSlaveCollector(url string, timeout time.Duration) *metricCollector {
	metrics := map[prometheus.Collector]func(metricMap, prometheus.Collector) error{}

	return newMetricCollector(url, timeout, metrics)
}