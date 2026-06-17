package clockseg

import (
	"sort"
)

// SortTurningPoints sorts the turning points in place by Index and Lambda.
func SortTurningPoints(tps []ClockTurningPoint) {
	sort.Slice(tps, func(i, j int) bool {
		if tps[i].Index != tps[j].Index {
			return tps[i].Index < tps[j].Index
		}
		return tps[i].Lambda < tps[j].Lambda
	})
}

// AreTurningPointsSorted checks if turning points are strictly ordered by index and lambda.
func AreTurningPointsSorted(tps []ClockTurningPoint) bool {
	for i := 1; i < len(tps); i++ {
		if tps[i].Index < tps[i-1].Index {
			return false
		}
		if tps[i].Index == tps[i-1].Index && tps[i].Lambda <= tps[i-1].Lambda {
			return false
		}
	}
	return true
}
