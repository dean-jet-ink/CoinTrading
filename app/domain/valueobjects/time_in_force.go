package valueobjects

import (
	"cointrading/app/domain/errors"
	"fmt"
	"net/http"
)

var (
	gtc = &TimeInForce{1}
	iok = &TimeInForce{2}
	fak = &TimeInForce{3}
	fas = &TimeInForce{4}
	fok = &TimeInForce{5}
	sok = &TimeInForce{6}
)

var timeInForces = map[int]string{
	1: "GTC",
	2: "IOK",
	3: "FAK",
	4: "FAS",
	5: "FOK",
	6: "SOK",
}

type TimeInForce struct {
	value int
}

func NewTimeInForce(value int) (*TimeInForce, error) {
	if _, ok := timeInForces[value]; !ok {
		message := fmt.Sprint("非対応の執行数量条件です")
		original := fmt.Sprintf("Unexpected time in force code %v", value)

		myerr := errors.NewMyError(message, original, http.StatusInternalServerError)

		return nil, myerr
	}

	return &TimeInForce{
		value: value,
	}, nil
}

func (t *TimeInForce) Value() int {
	return t.value
}

func (t *TimeInForce) DisplayValue() string {
	return timeInForces[t.value]
}

func (t *TimeInForce) IsGTC() bool {
	return *t == *gtc
}

func (t *TimeInForce) IsIOK() bool {
	return *t == *iok
}

func (t *TimeInForce) IsFAK() bool {
	return *t == *fak
}

func (t *TimeInForce) IsFAS() bool {
	return *t == *fas
}

func (t *TimeInForce) IsFOK() bool {
	return *t == *fok
}

func (t *TimeInForce) IsSOK() bool {
	return *t == *sok
}

func TimeInForces() []*TimeInForce {
	return []*TimeInForce{
		gtc,
		iok,
		fak,
		fas,
		fok,
		sok,
	}
}
