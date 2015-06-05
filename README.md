# go-timeslot-utils

Set of timeslot utilities for the Go programming language.

# Introduction

Sometimes, specially when working with measurement data, we may want to divide the data in timeslots. 
Timeslot are often used as partition keys on NO-SQL databases, like Cassandra.

This library only works with timeslot in hours, at the moment.

# Installation

```bash
go get "github.com/xavier-fernandez/go-timeslot-utils"
```

# Usage

### Import the library

```go
import (
	"github.com/xavier-fernandez/go-timeslot-utils"
)
```

#### OPTION 1 - Work with the timeslot in GMT

In case we are always working with multiple time zones, it makes sense working always using the Greenwich Meridian Time (GMT)

For obtaining the current timeslot we can call the following command:

```go
timeslot.CurrentTimeTimeslotInGMT()
```

For obtaining the timeslot of a certain time, the following command can be called:

```go
timeslot.TimeTimeslotInGMT(timestamp time.Time)
```

Both methods will return a timeslot string in the following format:

```go
"2015-06-03 07:00:00"
```
#### OPTION 2 - Work with the timeslot of a given Location

For obtaining the current timeslot we can call the following command:

```go
timeslot.CurrentTimeTimeslot(location time.Location)
```

For obtaining the timeslot of a certain time, the following command can be called:

```go
timeslot.TimeTimeslotInGMT(timestamp time.Time, location time.Location)
```

Both methods will return a timeslot string in the following format:

```go
"2015-05-02 22:00:00 -0400 EDT"
```
