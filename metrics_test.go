package metrics

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

const osTimerFreq = 1_000_000_000

func osTimer() int64 {
	return time.Now().UnixNano()
}

func TestOSTimer(t *testing.T) {
	runtime.LockOSThread() // Important: do not let scheduler switch us out
	defer runtime.UnlockOSThread()

	osStart := osTimer()
	var osEnd int64
	var osElapsed int64

	for osElapsed < osTimerFreq {
		osEnd = osTimer()
		osElapsed = osEnd - osStart
	}

	fmt.Printf("OS timer: %d -> %d = %d elapsed\n", osStart, osEnd, osElapsed)
	fmt.Printf("OS seconds: %.4f\n", float64(osElapsed)/float64(osTimerFreq))
}

func TestCPUTimer(t *testing.T) {
	runtime.LockOSThread() // Important: do not let scheduler switch us out
	defer runtime.UnlockOSThread()

	cpuStart := cpuTimer()
	osStart := osTimer()
	var osEnd int64
	var osElapsed int64

	for osElapsed < osTimerFreq {
		osEnd = osTimer()
		osElapsed = osEnd - osStart
	}

	cpuEnd := cpuTimer()
	cpuElapsed := cpuEnd - cpuStart

	fmt.Printf("OS timer: %d -> %d = %d elapsed\n", osStart, osEnd, osElapsed)
	fmt.Printf("OS seconds: %.4f\n", float64(osElapsed)/float64(osTimerFreq))
	fmt.Printf("CPU timer: %d -> %d = %d elapsed\n", cpuStart, cpuEnd, cpuElapsed)
}
