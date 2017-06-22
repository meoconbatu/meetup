package meetup

import "time"

const testVersion = 3

type WeekSchedule int

const (
	First WeekSchedule = iota + 1
	Second
	Third
	Fourth
	Last
	Teenth
)

func Day(week WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	x := 0
	if week == Last {
		month++
		week = First
		x = 7
	}
	d := time.Date(year, month, (int(week)*7)+1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, -int(time.Date(year, month, 8-int(weekday)-1, 0, 0, 0, 0, time.UTC).Weekday())-1)
	return d.AddDate(0, 0, -x).Day()
}
