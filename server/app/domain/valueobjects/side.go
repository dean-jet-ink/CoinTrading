package valueobjects

import (
	"cointrading/app/domain/myerror"
	"fmt"
	"strings"
)

type Side struct {
	value string
}

func NewSide(value string) (*Side, error) {
	lowerValue := strings.ToLower(value)

	if lowerValue != "buy" && lowerValue != "sell" {
		err := fmt.Errorf("%w: Invalid order side: %v", myerror.ErrInvalidOrderSide, lowerValue)

		return nil, err
	}

	return &Side{
		value: lowerValue,
	}, nil
}

func (s *Side) Value() string {
	return s.value
}

func (s *Side) DisplayValue() string {
	return strings.ToUpper(s.value)
}
