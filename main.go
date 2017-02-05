package main

import (
	"fmt"
	"net/url"
	"time"
)

func main() {
	fmt.Println("blah")

	var w Watch

	w.Gates = []int{1, 10, 100, 500, 1000}                  //alert me when each of these gates pass
	w.TTL = 5 * time.Second                                 //where the events happened in the last 30 minutes
	w.Address, _ = url.Parse("http://requestb.in/18hrvmi1") //by posting to this addresss
	w.Description = "500 returned to GPS"

	// a bad thing has happened.
	w.Fire()

	// another bad thing happens but no notify
	w.Fire()

	for i := 0; i < 2800; i++ {
		w.Fire() //the 10th event should should also generate a notification
	}

	// now we sleep 31 seconds and watch the watch clear
	// it should say All Clear: in the last 30 seconds there have been on 500 returned to GPS
	time.Sleep(10 * time.Second)

}
