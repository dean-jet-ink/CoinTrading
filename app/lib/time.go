package lib

import "time"

func StringToDateTime(timeStr string) (time.Time, error) {
	dateTime, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return time.Time{}, err
	}

	return dateTime, nil
}
