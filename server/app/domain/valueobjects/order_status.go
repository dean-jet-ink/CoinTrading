package valueobjects

import (
	"cointrading/app/domain/myerror"
	"fmt"
)

var (
	// 有効、または一部約定
	Ordered = &OrderStatus{1}
	// 全量約定
	Completed = &OrderStatus{2}
	// 約定待ち
	Waiting = &OrderStatus{3}
	// キャンセル
	Canceled = &OrderStatus{4}
	// 期限切れ
	Expired = &OrderStatus{5}
	// 注文失敗
	Rejected = &OrderStatus{6}
)

var statusList = map[OrderStatus]string{
	*Ordered:   "Ordered",
	*Completed: "Completed",
	*Waiting:   "Wainting",
	*Canceled:  "Canceled",
	*Expired:   "Expired",
	*Rejected:  "Rejected",
}

type OrderStatus struct {
	value int
}

func NewOrderStatus(value int) (*OrderStatus, error) {
	orderStatus := &OrderStatus{
		value: value,
	}

	if _, ok := statusList[*orderStatus]; !ok {
		err := fmt.Errorf("%w: Unexpected order status code: %v", myerror.ErrUnexpectedOrderStatus, value)

		return nil, err
	}

	return orderStatus, nil
}

func (o *OrderStatus) Value() int {
	return o.value
}

func (o *OrderStatus) DisplayValue() string {
	return statusList[*o]
}

func (o *OrderStatus) IsOrdered() bool {
	return *o == *Ordered
}

func (o *OrderStatus) IsCompleted() bool {
	return *o == *Completed
}

func (o *OrderStatus) IsWaiting() bool {
	return *o == *Waiting
}

func (o *OrderStatus) IsCanceled() bool {
	return *o == *Canceled
}

func (o *OrderStatus) IsExpired() bool {
	return *o == *Expired
}

func (o *OrderStatus) IsRejected() bool {
	return *o == *Rejected
}
