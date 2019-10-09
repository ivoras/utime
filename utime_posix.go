// +build darwin dragonfly freebsd linux netbsd openbsd

package utime

// #cgo CFLAGS: -g -Wall
// #include <time.h>
import "C"
import "time"

// Now returns a Time object initialised to the current time.
func Now() Time {
	// return Time(C.time(nil))
	// Go's runtime does magic which makes it much faster than the time(2) syscall
	return Time(time.Now().Unix())
}
