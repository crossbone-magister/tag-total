package output

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func sortTotalsKeyByTotalsValuesDescending(totals map[string]time.Duration) []string {
	keys := make([]string, 0, len(totals))
	for key := range totals {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return totals[keys[i]] > totals[keys[j]]
	})
	return keys
}

func FormatTotals(totals map[string]time.Duration) []string {
	longestKey := 0
	longestHour := 0
	for key, value := range totals {
		if len(key) > longestKey {
			longestKey = len(key)
		}
		durationString := value.String()
		if strings.ContainsRune(durationString, 'h') {
			hourLength := len(strings.Split(durationString, "h")[0])
			if hourLength > longestHour {
				longestHour = hourLength
			}
		}
	}

	sortedKeys := sortTotalsKeyByTotalsValuesDescending(totals)

	var output []string
	for _, key := range sortedKeys {
		r := strings.NewReplacer("h", " ", "m", " ", "s", " ")
		hms := strings.Fields(r.Replace(totals[key].String()))
		for len(hms) < 3 {
			hms = append([]string{"0"}, hms...)
		}
		formattedDuration := fmt.Sprintf("%*sh %2sm %2ss", longestHour, hms[0], hms[1], hms[2])
		output = append(output, fmt.Sprintf("%-*s - %s", longestKey, key, formattedDuration))
	}
	return output
}
