package service

import (
	"epictectus/appcontext"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
)

const (
	httpSubsystem   = "http"
	RequestCounter  = "requests_total"
	RequestLatency  = "request_duration_seconds"
	RequestInFlight = "requests_in_flight"
	unknownPath     = "unknown"
)

type Metric string
type TrackerInfo struct {
	Help   string
	Labels []string
}
type MuxRouteExtractor func(request *http.Request) string
type ServerMetricsPromCollector struct {
	Collectors            map[string]prometheus.Collector
	PathLabelFunc         MuxRouteExtractor
	HeaderLabels          []string
	requestLatencyBuckets []float64
}

var metricInfoMap map[Metric]TrackerInfo

const ApiRequestCount Metric = "api_request_count"
const ApiRequestDurationSeconds Metric = "api_request_duration_seconds"

func init() {
	metricInfoMap = make(map[Metric]TrackerInfo)
	metricInfoMap[ApiRequestCount] = TrackerInfo{
		Help:   "Api request count",
		Labels: []string{"endpoint", "status"},
	}
	metricInfoMap[ApiRequestDurationSeconds] = TrackerInfo{
		Help:   "Duration seconds for the API request",
		Labels: []string{"endpoint"},
	}
}
func NewCollector(options ...Option) *Collector {
	c := &Collector{Collectors: make(map[string]prometheus.Collector)}
	for _, opt := range options {
		opt(c)
	}
	c.initCollectors()
	return c
}

func CollectMetrics() {
	collector := appcontext.GetPrometheusHttpCollector()
	if collector != nil {
		collector.WithLabelValues().Add(1.0)
	}
}
