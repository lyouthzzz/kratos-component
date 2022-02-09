package metricsutil

import (
	"github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	"github.com/go-kratos/kratos/v2/metrics"
	goprometheus "github.com/prometheus/client_golang/prometheus"
	"os"
	"regexp"
)

var _appName = regexp.MustCompile("[.-]").ReplaceAllString(os.Getenv("APPLICATION_NAME"), "_")

func InitAppName(appName string) {
	_appName = appName
}

type (
	CounterVecOpts struct {
		Subsystem string
		Name      string
		Help      string
		Labels    []string
	}
	HistogramVecOpts struct {
		Subsystem string
		Name      string
		Help      string
		Labels    []string
		Buckets   []float64
	}
	GaugeVecOpts struct {
		Subsystem string
		Name      string
		Help      string
		Labels    []string
	}
)

func (opt CounterVecOpts) Build() metrics.Counter {
	vec := goprometheus.NewCounterVec(goprometheus.CounterOpts{
		Namespace: _appName,
		Subsystem: opt.Subsystem,
		Name:      opt.Name,
		Help:      opt.Help,
	}, opt.Labels)
	goprometheus.MustRegister(vec)
	return prometheus.NewCounter(vec)
}

func (opt HistogramVecOpts) Build() metrics.Observer {
	vec := goprometheus.NewHistogramVec(goprometheus.HistogramOpts{
		Namespace: _appName,
		Subsystem: opt.Subsystem,
		Name:      opt.Name,
		Help:      opt.Help,
		Buckets:   opt.Buckets,
	}, opt.Labels)
	goprometheus.MustRegister(vec)
	return prometheus.NewHistogram(vec)
}

func (opt GaugeVecOpts) Build() metrics.Gauge {
	vec := goprometheus.NewGaugeVec(goprometheus.GaugeOpts{
		Namespace: _appName,
		Subsystem: opt.Subsystem,
		Name:      opt.Name,
		Help:      opt.Help,
	}, opt.Labels)
	goprometheus.MustRegister(vec)
	return prometheus.NewGauge(vec)
}
