// Timers are for when you want to do something once in the future - tickers are for when
// you want to do something repeatedly at regular intervals. Here’s an example of a ticker
// that ticks periodically until we stop it.

package main

import "time"
import "fmt"

func main() {
	// Tickers use a similar mechanism to timers: a channel that is sent values. Here we’ll use
	// the range builtin on the channel to iterate over the values as they arrive every 500ms.
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	go func() {
		for t := range ticker.C {
			fmt.Println("Tack at", t)
		}
	}()

	go func() {
		for t := range ticker.C {
			fmt.Println("Toe at", t)
		}
	}()

	// Tickers can be stopped like timers. Once a ticker is stopped it won’t receive any more
	// values on its channel. We’ll stop ours after 1500ms.
	time.Sleep(time.Millisecond * 3000)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

// When we run this program the ticker should tick 3 times before we stop it.
