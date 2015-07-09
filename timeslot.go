package timeslot

import "time"

var gmt_location, _ = time.LoadLocation("GMT")

// Obtains the string of current timeslot.
// The string will be returned with the following format: '2015-05-02 22:00:00 -0400 EDT'
func CurrentTimeTimeslot(location *time.Location) string {
	return TimeTimeslot(time.Now(), location)
}

// Obtains the string of current timeslot in the GMT.
// The string will be returned with the following format: '2015-06-03 07:00:00'
func CurrentTimeTimeslotInGMT() string {
	return TimeTimeslotInGMT(time.Now())
}

// Obtains the timeslot of a given time.
// The string will be returned will be returned with the following format: '2015-05-02 22:00:00 -0400 EDT'
func TimeTimeslot(timestamp time.Time, location *time.Location) string {
	var timeslotHours = timestamp.Truncate(time.Hour)
	var timeslotString = timeslotHours.In(location).String()
	return timeslotString
}

// Obtains the timeslot of a given time in the GMT.
// The string will be returned will be returned with the following format: '2015-06-03 07:00:00'
func TimeTimeslotInGMT(timestamp time.Time) string {
	return TimeTimeslot(timestamp, gmt_location)[0:19]
}
