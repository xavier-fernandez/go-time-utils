package timeslot

import (
	"time"
)

// Obtains the string of current timeslot.
// The string will be returned with the following format: '2015-06-03 07:00:00'
// This method will always return the Greenwich Mean Time (GMT)
func CurrentTimeTimeslot() string {
	return TimeTimeslot(time.Now())
}

// Obtains the timeslot of a given time.
// The string will be returned will be returned with the following format: '2015-06-03 07:00:00'
// This method will always return the Greenwich Mean Time (GMT)
func TimeTimeslot(now time.Time) string {
	var location, _ = time.LoadLocation("GMT")
	var timeslotHours = now.Truncate(time.Hour)
	var timeslotString = timeslotHours.In(location).String()
	return timeslotString[0:19]
}
