package valueobjects

import (
	"cointrading/app/domain/myerror"
	"fmt"
)

var (
	GTC = &TimeInForce{1}
	IOK = &TimeInForce{2}
	FAK = &TimeInForce{3}
	FAS = &TimeInForce{4}
	FOK = &TimeInForce{5}
	SOK = &TimeInForce{6}
)

var timeInForces = map[TimeInForce]string{
	*GTC: "GTC",
	*IOK: "IOK",
	*FAK: "FAK",
	*FAS: "FAS",
	*FOK: "FOK",
	*SOK: "SOK",
}

type TimeInForce struct {
	value int
}

func NewTimeInForce(value int) (*TimeInForce, error) {
	timeInForce := &TimeInForce{
		value: value,
	}

	if _, ok := timeInForces[*timeInForce]; !ok {
		err := fmt.Errorf("%w: Unexpected time in force code %v", myerror.ErrUnexpectedTimeInForce, value)

		return nil, err
	}

	return timeInForce, nil
}

func (t *TimeInForce) Value() int {
	return t.value
}

func (t *TimeInForce) DisplayValue() string {
	return timeInForces[*t]
}

func (t *TimeInForce) IsGTC() bool {
	return *t == *GTC
}

func (t *TimeInForce) IsIOK() bool {
	return *t == *IOK
}

func (t *TimeInForce) IsFAK() bool {
	return *t == *FAK
}

func (t *TimeInForce) IsFAS() bool {
	return *t == *FAS
}

func (t *TimeInForce) IsFOK() bool {
	return *t == *FOK
}

func (t *TimeInForce) IsSOK() bool {
	return *t == *SOK
}

func TimeInForces() []*TimeInForce {
	return []*TimeInForce{
		GTC,
		IOK,
		FAK,
		FAS,
		FOK,
		SOK,
	}
}
