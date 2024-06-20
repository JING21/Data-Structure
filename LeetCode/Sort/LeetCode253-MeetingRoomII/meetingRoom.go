package LeetCode253_MeetingRoomII

import "sort"

func minMeetingRooms(intervals [][]int) int {
	type meeting struct {
		date   int
		updown int
	}

	var meetings []meeting
	for _, v := range intervals {
		meetings = append(meetings, meeting{v[0], 1})
		meetings = append(meetings, meeting{v[1], -1})
	}
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i].date < meetings[j].date || meetings[i].date == meetings[j].date && meetings[i].updown < meetings[j].updown
	})

	current, most := 0, 0

	for _, v := range meetings {
		current += v.updown
		most = max(most, current)
	}

	return most
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
