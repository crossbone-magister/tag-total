package main

import (
	"fmt"
	"os"

	"github.com/crossbone-magister/tag-total/logic"
	"github.com/crossbone-magister/tag-total/output"
	"github.com/crossbone-magister/timewlib"
)

func main() {
	parsed, err := timewlib.Parse(os.Stdin)
	if err == nil {
		intervals, err := timewlib.Process(parsed.Intervals)
		if err == nil {
			totals := logic.CalculateTotals(intervals)
			for _, total := range output.FormatTotals(totals) {
				fmt.Println(total)
			}
		} else {
			printErrorAndExit(err)
		}
	} else {
		printErrorAndExit(err)
	}

}

func printErrorAndExit(err error) {
	fmt.Printf("Error while parsing: %s\n", err)
	os.Exit(1)
}
