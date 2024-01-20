package myerror

import "errors"

var (
	ErrInvalidOrderSide       = errors.New("無効なオーダーサイドです")
	ErrUnexpectedOrderStatus  = errors.New("非対応のオーダーステータスです")
	ErrUnexpectedOrderType    = errors.New("非対応のオーダータイプです")
	ErrUnexpectedCurrencyCode = errors.New("非対応の貨幣コードです")
	ErrUnexpectedExchange     = errors.New("非対応の取引所です")
	ErrUnexpectedSymbol       = errors.New("非対応のシンボルです")
	ErrUnexpectedTimeInForce  = errors.New("非対応の執行数量条件です")
	ErrUnexpectedDuration     = errors.New("非対応の間隔です")
	ErrNotFoundOrder          = errors.New("オーダーが見つかりませんでした")
	ErrFailedToConnectNetwork = errors.New("ネットワーク接続に失敗しました")
	ErrBadRequest             = errors.New("不正なリクエストです")
	ErrRecordNotFound         = errors.New("レコードが見つかりませんでした")
)
