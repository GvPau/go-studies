package metrics

import (
	"sync"
	"time"
)

// Metrics struct holds our application metrics
type Metrics struct {
	sync.Mutex
	TotalRequests   int64
	RequestDuration []time.Duration
	RAMUsage        []uint64
}

// AppMetrics in an instance of Metrics to hold the application metrics
var AppMetrics = &Metrics{
	RequestDuration: make([]time.Duration, 0),
	RAMUsage:        make([]uint64, 0),
}
