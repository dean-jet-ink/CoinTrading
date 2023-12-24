package valueobjects

import (
	"cointrading/app/domain/errors"
	"fmt"
	"net/http"
	"strings"
)

type Side struct {
	value string
}

func NewSide(value string) (*Side, error) {
	lowerValue := strings.ToLower(value)

	if lowerValue != "buy" && lowerValue != "sell" {
		message := "無効なオーダーサイドです"
		original := fmt.Sprintf("Invalid order side: %v", lowerValue)
		myerr := errors.NewMyError(message, original, http.StatusInternalServerError)

		return nil, myerr
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
