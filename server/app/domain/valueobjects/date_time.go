package valueobjects

import (
	"time"
)

type DateTime struct {
	value time.Time
}

func NewDateTime(value time.Time) *DateTime {
	return &DateTime{
		value: value,
	}
}

func (d *DateTime) Value() time.Time {
	return d.value
}

func (d *DateTime) TruncateDateTime(duration time.Duration) time.Time {
	if duration == Month.Value() {
		return d.truncateMonth()
	}

	return d.value.Truncate(duration)
}

func (d *DateTime) truncateMonth() time.Time {
	y, m, _ := d.value.Date()

	return time.Date(y, m, 1, 0, 0, 0, 0, d.value.Location())
}
