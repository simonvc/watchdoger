package main

import (
	"net/url"
	"time"

	"github.com/simonvc/watchdoger"
)

func main() {
	var w watchdoger.Watch

	w.Gates = []int{10} //alert me when each of these gates pass
	//w.Gates = []int{1, 10, 100, 500, 1000}                                          //alert me when each of these gates pass
	w.TTL = 5 * time.Second                                                         //where the events happened in the last 30 minutes
	w.Address, _ = url.Parse("https://hooks.slack.com/services/webhookrulgoeshere") //by posting to this addresss
	w.Description = "An Error was returned to GPS"

	// a bad thing has happened.
	w.Fire()

	// another bad thing happens but no notify
	w.Fire()

	// then a few happen
	for i := 0; i < 20; i++ {
		w.Fire()
	}
	time.Sleep(10 * time.Second)
	// and then allclear

	// this shouldn't alert
	for i := 0; i < 5; i++ {
		w.Fire()
	}

	// and shoudln't allclear
	time.Sleep(10 * time.Second)

}
