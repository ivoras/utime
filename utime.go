package utime

// #cgo CFLAGS: -g -Wall
// #include <time.h>
import "C"

// Time is an efficient 32-bit timestamp type, counting the second from the Unix epoch.
// It doesn't (and can not) suport sub-second accuracy.
// utime.Time implements most of the API of Go's default time.Time type. It is always
// assumed to be in the UTC timezone, but there are functions to convert it to proper
// time.Time values with timezone information.
type Time uint32

// Duration records a duration in seconds
type Duration uint32

// Second is a Duration of 1 second
const Second Duration = 1

// Minute is a Duration of 1 minute
const Minute Duration = 60

// Hour is a Duration of 1 hour
const Hour Duration = 3600

// Day is a Duration of 1 day
const Day Duration = (3600 * 24)
