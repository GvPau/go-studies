package metrics

import (
	"net/http"
	"studies-1/internals/validation"
	"time"
)

type MetricsResponse struct {
	TotalRequests  int64   `json:"total_requests"`
	AvgRequestTime float64 `json:"avg_request_time_ms"`
	AvgRAMUsage    uint64  `json:"avg_ram_usage_bytes"`
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	AppMetrics.Lock()
	defer AppMetrics.Unlock()

	totalRequests := AppMetrics.TotalRequests
	var totalDuration time.Duration
	var totalRAMUsage uint64

	for _, duration := range AppMetrics.RequestDuration {
		totalDuration += duration
	}

	for _, ramUsage := range AppMetrics.RAMUsage {
		totalRAMUsage += ramUsage
	}

	avgRequestTime := float64(totalDuration.Microseconds()) / float64(totalRequests)
	avgRAMUsage := totalRAMUsage / uint64(totalRequests)

	response := MetricsResponse{
		TotalRequests:  totalRequests,
		AvgRequestTime: avgRequestTime,
		AvgRAMUsage:    avgRAMUsage,
	}

	validation.WriteJSONResponse(w, http.StatusOK, response)
}
