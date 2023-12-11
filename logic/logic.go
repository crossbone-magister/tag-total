package logic

import (
	"time"

	"github.com/crossbone-magister/timewlib"
)

func CalculateTotals(intervals []timewlib.Interval) map[string]time.Duration {
	totals := map[string]time.Duration{}
	for _, interval := range intervals {
		for _, tag := range interval.Tags {
			if totalDuration, ok := totals[tag]; ok {
				totals[tag] = totalDuration + interval.Duration()
			} else {
				totals[tag] = interval.Duration()
			}
		}
	}
	return totals
}
