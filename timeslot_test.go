package timeslot

import (
	"testing"
	"time"
)

const testNanoSecond = 5000
const testSecond = 21
const testMinute = 50
const testHour = 2
const testDay = 3
const testMonth = 5
const testYear = 2015

const testLocationGMT = "GMT"
const testLocationNY = "america/new_york"

const testTimeslotInGMT = "2015-05-03 02:00:00"
const testTimeslotNY = "2015-05-02 22:00:00 -0400 EDT"

//
// Returns a time.Time object for testing.
//
func getTestTime() time.Time {
	var location, _ = time.LoadLocation(testLocationGMT)
	return time.Date(testYear, testMonth, testDay, testHour, testMinute, testSecond, testNanoSecond, location)
}

//
// Test if all the preconditions needed for this class testings are generated as expected.
//
func TestTestPreconditions(t *testing.T) {
	// Test if the getTestTime method values are valid.
	var testTime = getTestTime()
	if testTime.Nanosecond() != testNanoSecond || testTime.Second() != testSecond ||
		testTime.Minute() != testMinute || testTime.Hour() != testHour ||
		testTime.Day() != testDay || testTime.Month() != testMonth ||
		testTime.Year() != testYear || testTime.Location().String() != testLocationGMT {
		t.Errorf("TestPreconditions -> The method getTestTime does not return a time object with valid attributes")
		t.Fail()
	}
}

//
// Test if CurrentTimeslot returns a string with a valid size.
//
func TestCurrentTimeslotInGMTSize(t *testing.T) {
	if len(CurrentTimeTimeslotInGMT()) != len(testTimeslotInGMT) {
		t.Errorf("TestCurrentTimeslotInGMTSize -> The method TestCurrentTimeslotInGMT does not return a valid lenght")
		t.Fail()
	}
}

//
// Test if the TestTimeTimeslot returns a valid value.
//
func TestTimeTimeslotInGMTValue(t *testing.T) {
	var timeTimeslot = TimeTimeslotInGMT(getTestTime())
	if timeTimeslot != testTimeslotInGMT {
		t.Errorf("TestTimeTimeslotInGMTValue -> The timeslot '", timeTimeslot, "' does not correspond to the timeslot: ", testTimeslotInGMT)
		t.Fail()
	}
}

//
// Test if CurrentTimeslot returns a string with a valid size.
//
func TestCurrentTimeslotSize(t *testing.T) {
	var location, _ = time.LoadLocation(testLocationNY)
	if len(CurrentTimeTimeslot(location)) != len(testTimeslotNY) {
		t.Errorf("TestCurrentTimeslotSize -> The method CurrentTimeslot does not return a valid lenght")
		t.Fail()
	}
}

//
// Test if the TestTimeTimeslot returns a valid value.
//
func TestTimeTimeslotValue(t *testing.T) {
	var location, _ = time.LoadLocation(testLocationNY)
	var testTime = getTestTime()
	var timeTimeslot = TimeTimeslot(testTime, location)
	if timeTimeslot != testTimeslotNY {
		t.Errorf("TestTimeTimeslot -> The timeslot '", timeTimeslot, "' does not correspond to the timeslot: ", testTimeslotNY)
		t.Fail()
	}
}

//
// Benchmarks the CurrentTimeslotInGMT method.
//
func BenchmarkCurrentTimeslotInGMT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CurrentTimeTimeslotInGMT()
	}
}

//
// Benchmarks the TimeTimeslotInGMT method.
//
func BenchmarkTimeTimeslotInGMT(b *testing.B) {
	var time = getTestTime()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TimeTimeslotInGMT(time)
	}
}

//
// Benchmarks the CurrentTimeslot method.
//
func BenchmarkCurrentTimeslot(b *testing.B) {
	var location, _ = time.LoadLocation(testLocationNY)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CurrentTimeTimeslot(location)
	}
}

//
// Benchmarks the TimeTimeslot method.
//
func BenchmarkTimeTimeslot(b *testing.B) {
	var testTime = getTestTime()
	var location, _ = time.LoadLocation(testLocationNY)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TimeTimeslot(testTime, location)
	}
}
