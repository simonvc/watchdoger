package main

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/hpcloud/tail"
	"github.com/simonvc/watchdoger"
)

// watches an nginx log file for status code 499, but only lets us know if more than 1000 happen in 30 mins.

var (
	fourNinetyNineSeen watchdoger.Watch
	totalRequests      int
	total499           int
)

func main() {
	fourNinetyNineSeen.Gates = []int{1000, 2000, 3000, 4000, 5000, 10000}
	fourNinetyNineSeen.TTL = 30 * time.Minute
	fourNinetyNineSeen.Address, _ = url.Parse("https://hooks.slack.com/services/your/slack/webhook")
	fourNinetyNineSeen.Description = "499 in log file. "

	t, _ := tail.TailFile(os.Args[1], tail.Config{Follow: true})
	for line := range t.Lines {
		totalRequests++
		if strings.Contains(line.Text, "499") {
			total499++
			fmt.Printf("499 seen. Count in last 30 mins: %d   Total499/Total: %d/%d = %f %% \n",
				fourNinetyNineSeen.Current, total499, totalRequests, float64(total499)/float64(totalRequests))
			fourNinetyNineSeen.Fire()
		}

	}
}
