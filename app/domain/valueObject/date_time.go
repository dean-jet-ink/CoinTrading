package valueobject

import "time"

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
	return d.value.Truncate(duration)
}
