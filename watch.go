package main

import (
	"fmt"
	"math"
	"net/url"
	"sync"
	"time"
)

// Watch a simple way to notify on errors that's less spammy than
// sending a webhoook for every event.
type Watch struct {
	lock        sync.Mutex
	Gates       []int
	TTL         time.Duration
	Address     *url.URL
	Current     int
	Description string
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func humanizeDuration(duration time.Duration) string {
	if duration.Seconds() < 60.0 {
		return fmt.Sprintf("%d seconds", int64(duration.Seconds()))
	}
	if duration.Minutes() < 60.0 {
		remainingSeconds := math.Mod(duration.Seconds(), 60)
		return fmt.Sprintf("%d minutes %d seconds", int64(duration.Minutes()), int64(remainingSeconds))
	}
	if duration.Hours() < 24.0 {
		remainingMinutes := math.Mod(duration.Minutes(), 60)
		remainingSeconds := math.Mod(duration.Seconds(), 60)
		return fmt.Sprintf("%d hours %d minutes %d seconds",
			int64(duration.Hours()), int64(remainingMinutes), int64(remainingSeconds))
	}
	remainingHours := math.Mod(duration.Hours(), 24)
	remainingMinutes := math.Mod(duration.Minutes(), 60)
	remainingSeconds := math.Mod(duration.Seconds(), 60)
	return fmt.Sprintf("%d days %d hours %d minutes %d seconds",
		int64(duration.Hours()/24), int64(remainingHours),
		int64(remainingMinutes), int64(remainingSeconds))
}

func (w *Watch) Fire() {
	//fmt.Println("i fired")
	w.lock.Lock()
	defer w.lock.Unlock()
	// increment the counter
	w.Current++
	// if the counter is in the list of gates
	// fire the webhook
	if contains(w.Gates, w.Current) {
		fmt.Printf("%d x Error: %s in the last %s\n", w.Current, w.Description, humanizeDuration(w.TTL))
	}
	// fire off a goroutine to reduce after w.ttl
	go comeback(w)
}

func (w *Watch) extinguish() {
	w.lock.Lock()
	defer w.lock.Unlock()
	// decrement the counter
	w.Current--
	// if the counter < 1 fire the webhook to say allclear
	if w.Current < 1 {
		fmt.Printf("All Clear: No %s seen in the last %s\n", w.Description, humanizeDuration(w.TTL))
	}
}

func comeback(w *Watch) {
	time.Sleep(w.TTL)
	w.extinguish()
}
