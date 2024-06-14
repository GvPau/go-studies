package middlewares

import (
	"net/http"
	"runtime"
	"studies-1/internals/metrics"
	"time"
)

func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Measure RAM usage before handling request
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		startRAM := m.TotalAlloc

		next.ServeHTTP(w, r)

		// Measure RAM usage after handling request
		runtime.ReadMemStats(&m)
		endRAM := m.TotalAlloc
		ramUsed := endRAM - startRAM

		// Update metrics with RAM usage
		duration := time.Since(start)

		metrics.AppMetrics.Lock()
		metrics.AppMetrics.TotalRequests++
		metrics.AppMetrics.RequestDuration = append(metrics.AppMetrics.RequestDuration, duration)
		metrics.AppMetrics.RAMUsage = append(metrics.AppMetrics.RAMUsage, ramUsed)
		metrics.AppMetrics.Unlock()

	})
}
