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
const testLocation = "GMT"
const testTimeslot = "2015-05-03 02:00:00"

//
// Returns a time.Time object for testing.
//
func getTestTime(t *testing.T) time.Time {
	var location, err = time.LoadLocation(testLocation)
	if err != nil {
		t.Errorf("getTestTime -> The following error was trown when instantiating the ", err.Error())
		t.Fail()
	}
	return time.Date(testYear, testMonth, testDay, testHour, testMinute, testSecond, testNanoSecond, location)
}

//
// Test if all the preconditions needed for this class testings are generated as expected.
//
func TestTestPreconditions(t *testing.T) {
	// Test if the getTestTime method values are valid.
	var testTime = getTestTime(t)
	if testTime.Nanosecond() != testNanoSecond || testTime.Second() != testSecond ||
		testTime.Minute() != testMinute || testTime.Hour() != testHour ||
		testTime.Day() != testDay || testTime.Month() != testMonth ||
		testTime.Year() != testYear || testTime.Location().String() != testLocation {
		t.Errorf("TestPreconditions -> The method getTestTime does not return a time object with valid attributes")
		t.Fail()
	}
}

//
// Test if CurrentTimeslot returns a string with a valid size.
//
func TestCurrentTimeslotSize(t *testing.T) {
	if len(CurrentTimeTimeslot()) != len("2015-06-03 07:00:00") {
		t.Errorf("TestCurrentTimeslotSize -> The method CurrentTimeslot does not return a valid lenght")
		t.Fail()
	}
}

//
// Test if the TestTimeTimeslot returns a valid value.
//
func TestTimeTimeslotValue(t *testing.T) {
	var timeTimeslot = TimeTimeslot(getTestTime(t))
	if timeTimeslot != "2015-05-03 02:00:00" {
		t.Errorf("TestTimeTimeslot -> The timeslot '", timeTimeslot, "' is not correspond to the timeslot: ", testTimeslot)
		t.Fail()
	}
}

//
// Benchmarks the CurrentTimeslot method.
//
func BenchmarkCurrentTimeslot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CurrentTimeTimeslot()
	}
}

//
// Benchmarks the TimeTimeslot method.
//
func BenchmarkTimeTimeslot(b *testing.B) {
	var time = getTestTime(nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TimeTimeslot(time)
	}
}
