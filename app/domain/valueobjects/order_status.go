package valueobjects

import (
	"cointrading/app/domain/errors"
	"fmt"
	"net/http"
)

var (
	// 有効、または一部約定
	ordered = &OrderStatus{1}
	// 全量約定
	completed = &OrderStatus{2}
	// 約定待ち
	waiting = &OrderStatus{3}
	// キャンセル
	canceled = &OrderStatus{4}
	// 期限切れ
	expired = &OrderStatus{5}
)

var statusList = map[int]string{
	1: "Ordered",
	2: "Completed",
	3: "Wainting",
	4: "Canceled",
	5: "Expired",
}

type OrderStatus struct {
	value int
}

func NewOrderStatus(value int) (*OrderStatus, error) {
	if _, ok := statusList[value]; !ok {
		message := fmt.Sprint("無効なオーダーステータスです")
		original := fmt.Sprintf("Invalid order status code: %v", value)

		myerr := errors.NewMyError(message, original, http.StatusInternalServerError)

		return nil, myerr
	}

	return &OrderStatus{
		value: value,
	}, nil
}

func (o *OrderStatus) Value() int {
	return o.value
}

func (o *OrderStatus) DisplayValue() string {
	return statusList[o.value]
}

func (o *OrderStatus) IsOrdered() bool {
	return *o == *ordered
}

func (o *OrderStatus) IsCompleted() bool {
	return *o == *completed
}

func (o *OrderStatus) IsWaiting() bool {
	return *o == *waiting
}

func (o *OrderStatus) IsCanceled() bool {
	return *o == *canceled
}

func (o *OrderStatus) IsExpired() bool {
	return *o == *expired
}
