package valueobjects

import (
	"cointrading/app/domain/myerror"
	"fmt"
	"strings"
	"time"
)

var (
	Second = &Duration{time.Second}
	Minute = &Duration{time.Minute}
	Hour   = &Duration{time.Hour}
	Day    = &Duration{time.Hour * 24}
	Week   = &Duration{time.Hour * 24 * 7}
	Month  = &Duration{time.Hour * 24 * 30}
)

var (
	durationNames = map[Duration]string{
		*Second: "Second",
		*Minute: "Minute",
		*Hour:   "Hour",
		*Day:    "Day",
		*Week:   "Week",
		*Month:  "Month",
	}
)

type Duration struct {
	value time.Duration
}

func NewDuration(value time.Duration) (*Duration, error) {
	duration := &Duration{
		value: value,
	}

	if _, ok := durationNames[*duration]; !ok {
		err := fmt.Errorf("%w: Unexpected duration code: %v", myerror.ErrUnexpectedDuration, value)

		return nil, err
	}

	return duration, nil
}

func (d *Duration) Value() time.Duration {
	return d.value
}

func (d *Duration) DisplayValue() string {
	return durationNames[*d]
}

func (d *Duration) DisplayValueForTableName() string {
	return strings.ToLower(durationNames[*d])
}

func (d *Duration) IsSecond() bool {
	return d.value == Second.Value()
}

func (d *Duration) IsMinute() bool {
	return d.value == Minute.Value()
}

func (d *Duration) IsHour() bool {
	return d.value == Hour.Value()
}

func (d *Duration) IsDay() bool {
	return d.value == Day.Value()
}

func (d *Duration) IsWeek() bool {
	return d.value == Week.Value()
}

func (d *Duration) IsMonth() bool {
	return d.value == Month.Value()
}

func Durations() []*Duration {
	return []*Duration{
		Second,
		Minute,
		Hour,
		Day,
		Week,
		Month,
	}
}
