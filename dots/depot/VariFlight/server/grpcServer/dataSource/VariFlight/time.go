// Scry Info.  All rights reserved.
// license that can be found in the license file.

package VariFlight

import (
	"strconv"
	"time"
)

// parseLocalToUTC parses the given local time value, and then returns the Time with the location set to UTC
func parseLocalToUTC(layout, value, timeZoneNameOrOffsetStr string) (time.Time, error) {
	location, err := parseLocation(timeZoneNameOrOffsetStr)
	if err != nil {
		return time.Time{}, err
	}

	t, err := time.ParseInLocation(layout, value, location)
	if err != nil {
		return time.Time{}, err
	}

	return t.UTC(), nil
}

// nowToUTC returns the current time with the location set to UTC.
func nowToUTC() time.Time {
	return time.Now().UTC()
}

// parseLocation returns the Location with the given timeZoneNameOrOffsetStr, which must either a location name corresponding to
// a file in the IANA Time Zone database, such as "Asia/Hong_Kong", or a fixed offset (seconds east of UTC) string, such as "28800".
func parseLocation(timeZoneNameOrOffsetStr string) (*time.Location, error) {
	if offset, err := strconv.Atoi(timeZoneNameOrOffsetStr); err != nil {
		return time.LoadLocation(timeZoneNameOrOffsetStr)
	} else {
		return time.FixedZone(timeZoneNameOrOffsetStr, offset), nil
	}
}
