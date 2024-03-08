// only used in rapidgator package.
// a simple limit rate implement.

package rapidgator

import (
	"math/rand"
	"sync"
	"time"
)

type LimitRateChecker struct {
	LastCheckTime time.Time
	Mux           sync.Mutex
}

var limitRateChecker = LimitRateChecker{}

// limit interval: Second.
const baseLimitRateInterval = 10
const randomLimitRateInterval = 5

func (limitRateChecker *LimitRateChecker) check() bool {
	limitRateChecker.Mux.Lock()
	defer limitRateChecker.Mux.Unlock()

	isLimited := false
	current := time.Now()

	// first check, return true.
	if limitRateChecker.LastCheckTime.IsZero() {
		limitRateChecker.LastCheckTime = current

		return isLimited
	}

	duration := current.Sub(limitRateChecker.LastCheckTime)
	checkInterval := time.Duration(baseLimitRateInterval+rand.Intn(randomLimitRateInterval)) * time.Second
	isLimited = duration <= checkInterval
	if !isLimited {
		limitRateChecker.LastCheckTime = current
	}

	return isLimited
}
