package ratelimit

import (
	"sync"
	"time"
)

type attempt struct {
	count     int
	blockedAt time.Time
	lastSeen  time.Time
}

const (
	MaxAttempts     = 10
	BlockDuration   = 15 * time.Minute
	cleanupInterval = 1 * time.Hour
)

var (
	mu       sync.Mutex
	attempts = make(map[string]*attempt)
)

func init() {
	go func() {
		for range time.Tick(cleanupInterval) {
			mu.Lock()
			cutoff := time.Now().Add(-BlockDuration * 2)
			for ip, a := range attempts {
				if a.lastSeen.Before(cutoff) {
					delete(attempts, ip)
				}
			}
			mu.Unlock()
		}
	}()
}

// RecordFailure records a failed login attempt from the given IP.
func RecordFailure(ip string) {
	mu.Lock()
	defer mu.Unlock()
	a, ok := attempts[ip]
	if !ok {
		a = &attempt{}
		attempts[ip] = a
	}
	a.count++
	a.lastSeen = time.Now()
	if a.count >= MaxAttempts {
		a.blockedAt = time.Now()
	}
}

// Reset clears the failure counter for an IP on successful login.
func Reset(ip string) {
	mu.Lock()
	defer mu.Unlock()
	delete(attempts, ip)
}

// IsBlocked returns true if the IP has exceeded the allowed attempts.
func IsBlocked(ip string) bool {
	mu.Lock()
	defer mu.Unlock()
	a, ok := attempts[ip]
	if !ok {
		return false
	}
	if a.count < MaxAttempts {
		return false
	}
	if time.Since(a.blockedAt) > BlockDuration {
		delete(attempts, ip)
		return false
	}
	return true
}
