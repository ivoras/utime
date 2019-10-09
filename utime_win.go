// +build windows

package utime

import "time"

const nanosecondsScale int64 = 1000000000

// Now returns a Time object initialised to the current time.
func Now() Time {
	/*
		ft := windows.Filetime{}
		windows.GetSystemTimeAsFileTime(&ft)
		return Time(ft.Nanoseconds() / nanosecondsScale)
	*/
	// On Windows, it turns out to be much faster to simply call time.Now().Unix(), so do it
	return Time(time.Now().Unix())
}
