package sports

import (
	"fmt"
	"time"
)

const msInHour = 60 * 60 * 1000

// SecondsToTimeString converts seconds to a string in the format "HH:MM:SS".
func SecondsToTimeString(seconds int64) string {
	hours := (seconds * 1000) / msInHour
	t := time.UnixMilli(seconds * 1000)

	// https://pkg.go.dev/time#Time.Format
	//
	// The layout string used by the Parse function and Format method
	// shows by example how the reference time should be represented.
	// We stress that one must show how the reference time is formatted,
	// not a time of the user's choosing. Thus each layout string is a
	// representation of the time stamp,
	//	Jan 2 15:04:05 2006 MST
	// An easy way to remember this value is that it holds, when presented
	// in this order, the values (lined up with the elements above):
	//	  1 2  3  4  5    6  -7
	return fmt.Sprintf("%02d:%s", hours, t.Format("04:05"))
}
