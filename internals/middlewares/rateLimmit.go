package middlewares

import (
	"log"
	"net/http"
	"sync"
	"time"
)

// A map that uses the client's IP as that key and the rateLimiter struct as the value
// Mutex is used to ensure that concurrent access to rateLimiter is save
var (
	rateLimitMap = make(map[string]*rateLimiter)
	mu           sync.Mutex
)

// A structure that holds the number of requests and the timestamp of the first request in the current window for each IP
type rateLimiter struct {
	requests  int
	timestamp time.Time
}

const (
	rateLimit       = 5               // requests
	rateLimitWindow = 1 * time.Minute // per minute
)

func RateLimitMiddlware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr

		mu.Lock()
		rl, exists := rateLimitMap[ip]
		if !exists || time.Since(rl.timestamp) > rateLimitWindow {
			log.Printf("New rate limit window for IP: %s", ip)
			rl = &rateLimiter{
				requests:  1,
				timestamp: time.Now(),
			}
			rateLimitMap[ip] = rl
		} else {
			rl.requests++
			log.Printf("IP: %s - Requests: %d", ip, rl.requests)

		}
		mu.Unlock()

		if rl.requests > rateLimit {
			log.Printf("IP: %s has exceeded the rate limit", ip)
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
