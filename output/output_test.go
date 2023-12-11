package output_test

import (
	"testing"
	"time"

	"github.com/crossbone-magister/tag-total/output"
)

func TestPrintTotals(t *testing.T) {
	tag1Duration, err := time.ParseDuration("30m")
	if err != nil {
		t.Fatalf("Error while parsing duration: %s", err)
	}
	tag2Duration, err := time.ParseDuration("2h")
	if err != nil {
		t.Fatalf("Error while parsing duration: %s", err)
	}
	tag3Duration, err := time.ParseDuration("1h30m")
	if err != nil {
		t.Fatalf("Error while parsing duration: %s", err)
	}
	tag4Duration, err := time.ParseDuration("30s")
	if err != nil {
		t.Fatalf("Error while parsing duration: %s", err)
	}
	totals := map[string]time.Duration{
		"tag1": tag1Duration,
		"tag2": tag2Duration,
		"tag3": tag3Duration,
		"tag4": tag4Duration,
	}
	output := output.FormatTotals(totals)
	expectedOutput := []string{"tag2 - 2h  0m  0s", "tag3 - 1h 30m  0s", "tag1 - 0h 30m  0s", "tag4 - 0h  0m 30s"}
	for i, expected := range expectedOutput {
		if expected != output[i] {
			t.Errorf("Output [%s] is different from expected output [%s]", output[i], expected)
		}
	}
}

func TestPrintTotalsFormats(t *testing.T) {
	tag1Duration, err := time.ParseDuration("20h30m")
	if err != nil {
		t.Fatalf("Error while parsing duration: %s", err)
	}
	tag2Duration, err := time.ParseDuration("2h5m")
	if err != nil {
		t.Fatalf("Error while parsing duration: %s", err)
	}
	tag3Duration, err := time.ParseDuration("1h30m40s")
	if err != nil {
		t.Fatalf("Error while parsing duration: %s", err)
	}
	totals := map[string]time.Duration{
		"tag1": tag1Duration,
		"tag2": tag2Duration,
		"tag3": tag3Duration,
	}
	output := output.FormatTotals(totals)
	expectedOutput := []string{"tag1 - 20h 30m  0s", "tag2 -  2h  5m  0s", "tag3 -  1h 30m 40s"}
	for i, expected := range expectedOutput {
		if expected != output[i] {
			t.Errorf("Output [%s] is different from expected output [%s]", output[i], expected)
		}
	}
}
