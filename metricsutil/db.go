package metricsutil

import (
	"github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	"github.com/go-kratos/kratos/v2/metrics"
	goprometheus "github.com/prometheus/client_golang/prometheus"
)

var (
	_metricSuccessTotal   *goprometheus.CounterVec
	_metricErrorTotal     *goprometheus.CounterVec
	_metricSuccessLatency *goprometheus.HistogramVec
	_metricErrorLatency   *goprometheus.HistogramVec

	DBSuccessTotal   metrics.Counter
	DBErrorTotal     metrics.Counter
	DBSuccessLatency metrics.Observer
	DBErrorLatency   metrics.Observer
)

func init() {
	_metricSuccessTotal = goprometheus.NewCounterVec(goprometheus.CounterOpts{
		Namespace: "db",
		Subsystem: "requests",
		Name:      "success_total",
		Help:      "The total number of db operation",
	}, []string{"system", "operation"})
	_metricErrorTotal = goprometheus.NewCounterVec(goprometheus.CounterOpts{
		Namespace: "db",
		Subsystem: "requests",
		Name:      "error_total",
		Help:      "The second latency of db operation",
	}, []string{"system", "operation"})
	_metricSuccessLatency = goprometheus.NewHistogramVec(goprometheus.HistogramOpts{
		Namespace: "db",
		Subsystem: "requests",
		Name:      "success_latency_seconds",
		Help:      "The second latency of db operation",
	}, []string{"system", "operation"})
	_metricErrorLatency = goprometheus.NewHistogramVec(goprometheus.HistogramOpts{
		Namespace: "db",
		Subsystem: "requests",
		Name:      "error_latency_seconds",
		Help:      "The second latency of db operation",
	}, []string{"system", "operation"})
	goprometheus.MustRegister(_metricSuccessTotal, _metricErrorTotal, _metricSuccessLatency, _metricErrorLatency)

	DBSuccessTotal = prometheus.NewCounter(_metricSuccessTotal)
	DBErrorTotal = prometheus.NewCounter(_metricErrorTotal)
	DBSuccessLatency = prometheus.NewHistogram(_metricSuccessLatency)
	DBErrorLatency = prometheus.NewHistogram(_metricErrorLatency)
}
