package logic_test

import (
	"testing"
	"time"

	"github.com/crossbone-magister/tag-total/logic"
	"github.com/crossbone-magister/timewlib"
)

func initExpectations(t *testing.T) map[string]time.Duration {
	tag1Duration, err := time.ParseDuration("1h")
	if err != nil {
		t.Fatalf("Error initializing duration: %s", err)
	}
	tag2Duration, err := time.ParseDuration("30m")
	if err != nil {
		t.Fatalf("Error initializing duration: %s", err)
	}
	return map[string]time.Duration{
		"tag1": tag1Duration,
		"tag2": tag2Duration,
	}
}

func TestCalculateTotals(t *testing.T) {
	expectations := initExpectations(t)
	first := timewlib.NewInterval(10, 0, 10, 30)
	first.Tags = []string{"tag1"}
	second := timewlib.NewInterval(10, 30, 11, 00)
	second.Tags = []string{"tag2"}
	third := timewlib.NewInterval(11, 00, 11, 30)
	third.Tags = []string{"tag1"}
	intervals := []timewlib.Interval{
		*first,
		*second,
		*third,
	}
	totals := logic.CalculateTotals(intervals)
	for tag, expectedDuration := range expectations {
		actualDuration := totals[tag]
		if actualDuration != expectedDuration {
			t.Errorf("Duration for tag %s is not %s, but %s", tag, expectedDuration, actualDuration)
		}
	}
}
