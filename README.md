# Watchdoger

A helpful pattern to send a notification to slack when something bad happens, in a non spammy way.

Checkout example/main.go to see how to use it.

Import it

```
import "github.com/simonvc/watchdoger"
```


Make a global Watch

```
var w watchdoger.Watch
```

Configure it with the number of errors that should trigger an alert, and over what time period.

````
    w.Gates = []int{1, 10, 100, 500, 1000}                                                                    //alert me when each of these gates pass
    w.TTL = 15 * time.Second                                                                                   //where the events happened in the last 30 minutes
    w.Address, _ = url.Parse("https://hooks.slack.com/services/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") //by posting to this addresss
    w.Description = "An Error occured!"
```


Now, when something goes wrong, it'll alert you, but only if you've passed the threshold in the last N minutes.

```
    // a single bad thing has happened.
    w.Fire()

    // another bad thing happens but don't notify me
    w.Fire()

    // oh shit, lots of bad  things are happening.
    for i := 0; i < 2800; i++ {
        w.Fire() //the 10th event should should also generate a notification
    }
```

Finally, when no errors have occurerd in the window, you'll get an all clear.

```
    // now we sleep 31 seconds and watch the watch clear
    // it should say All Clear: in the last 30 seconds there have been no errors
    time.Sleep(31 * time.Second
```
