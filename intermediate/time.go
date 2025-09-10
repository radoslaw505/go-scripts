package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Current time:", now)

	// Formatting
	fmt.Println("Formatted:", now.Format("2006-01-02 15:04:05"))

	// Parsing
	layout := "2006-01-02"
	parsed, err := time.Parse(layout, "2025-09-09")
	if err != nil {
		fmt.Println("Parse error:", err)
	} else {
		fmt.Println("Parsed date:", parsed)
	}

	// Duration
	twoHoursLater := now.Add(2 * time.Hour)
	fmt.Println("2 hours later:", twoHoursLater)

	// Comparison
	future := now.Add(5 * time.Minute)
	fmt.Println("Is now before future?", now.Before(future))

	// Time zone
	warsaw, err := time.LoadLocation("Europe/Warsaw")
	if err == nil {
		fmt.Println("Time in Warsaw:", now.In(warsaw))
	}

	// Sleep
	fmt.Println("Sleeping for 1 second...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done sleeping.")

	// Rounding
	rounded := now.Round(time.Minute)
	fmt.Println("Rounded to nearest minute:", rounded)

	// Ticker
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Ticker ticked at:", t)
		}
	}()
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped.")

	// Timer
	timer := time.NewTimer(5 * time.Second)
	<-timer.C
	fmt.Println("Timer expired after 5 seconds.")

	// Unix time
	fmt.Println("Unix time:", now.Unix(), "seconds since epoch")
	fmt.Println("Unix nano time:", now.UnixNano(), "nanoseconds since epoch")
	fmt.Println("Unix milli time:", now.UnixMilli(), "milliseconds since epoch")

	// Convert Unix timestamp to time.Time
	ts := int64(1694280000)
	converted := time.Unix(ts, 0)
	fmt.Println("Converted from epoch:", converted)

	// Convert back to epoch
	fmt.Println("Back to Unix:", converted.Unix())
}
