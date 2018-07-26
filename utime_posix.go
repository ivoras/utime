// +build darwin dragonfly freebsd linux netbsd openbsd

package utime

// #cgo CFLAGS: -g -Wall
// #include <time.h>
import "C"

// Now returns a Time object initialised to the current time.
func Now() Time {
	return Time(C.time(nil))
}
