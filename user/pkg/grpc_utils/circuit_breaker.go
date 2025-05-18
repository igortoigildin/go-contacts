package grpcutils

import (
	"log"
	"time"

	"github.com/sony/gobreaker"
)

func InitCircuitBreaker() *gobreaker.CircuitBreaker {
	settings := gobreaker.Settings{
		Name:        "MyGRPCCircuitBreaker",
		MaxRequests: 3,               // Allowed requests in half-open state
		Interval:    0,               // No rolling window; state resets manually/timed
		Timeout:     5 * time.Second, // Time to wait before going from open to half-open
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			// Open breaker after 5 failures
			return counts.ConsecutiveFailures >= 5
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			log.Printf("Circuit breaker state changed from %s to %s\n", from.String(), to.String())
		},
	}

	return gobreaker.NewCircuitBreaker(settings)
}
