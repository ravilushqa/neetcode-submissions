import (
	"slices"
	"cmp"
)

/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func canAttendMeetings(intervals []Interval) bool {
	slices.SortFunc(intervals, func(a, b Interval) int {
		return cmp.Compare(a.start, b.start)
	})

	for i := range intervals {
		if i == 0 {
			continue
		}

		if intervals[i].start < intervals[i-1].end {
			return false
		} 
	}

	return true
}
