package bitflyer

import (
	"cointrading/app/domain/errors"
	"cointrading/app/domain/valueobjects"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

func (b *bitflyerClient) privateAPIHeader(method, path string, body []byte) (map[string]string, error) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	message := timestamp + method + path + string(body)

	h := hmac.New(sha256.New, []byte(b.secret))

	if _, err := h.Write([]byte(message)); err != nil {
		return nil, err
	}

	sign := hex.EncodeToString(h.Sum(nil))

	header := map[string]string{
		"ACCESS-KEY":       b.key,
		"ACCESS-TIMESTAMP": timestamp,
		"ACCESS-SIGN":      sign,
		"Content-Type":     "application/json",
	}

	return header, nil
}

func (b *bitflyerClient) publicAPIHeader() map[string]string {
	header := map[string]string{
		"Content-Type": "application/json",
	}

	return header
}

func (b *bitflyerClient) convertSymbol(symbol *valueobjects.Symbol) (string, error) {
	switch {
	case symbol.IsBTCJPY():
		return "BTC_JPY", nil
	case symbol.IsETHJPY():
		return "ETH_JPY", nil
	default:
		message := "非対応のシンボルです"
		original := fmt.Sprintf("Unexpected symbol code %v", symbol.Value())

		myerr := errors.NewMyError(message, original, 500)

		return "", myerr
	}
}

func (b *bitflyerClient) convertCurrencyCode(currencyCode string) (int, bool) {
	currencyCodes := map[string]int{
		"JPY": valueobjects.JPY.Value(),
		"BTC": valueobjects.BTC.Value(),
		"ETH": valueobjects.ETH.Value(),
	}

	code, ok := currencyCodes[currencyCode]
	if !ok {
		return 0, false
	}

	return code, true
}

func (b *bitflyerClient) convertOrderType(orderType *valueobjects.OrderType) (string, error) {
	switch {
	case orderType.IsMarket():
		return "MARKET", nil
	case orderType.IsLimit():
		return "LIMIT", nil
	case orderType.IsStop():
		return "STOP", nil
	default:
		message := "非対応のオーダータイプです"
		original := fmt.Sprintf("Unexpected order type code %v", orderType.Value())

		myerr := errors.NewMyError(message, original, 500)

		return "", myerr
	}
}

func (b *bitflyerClient) convertTimeInForce(timeInForce *valueobjects.TimeInForce) (string, error) {
	switch {
	case timeInForce.IsGTC():
		return "GTC", nil
	case timeInForce.IsIOK():
		return "IOK", nil
	case timeInForce.IsFAK():
		return "FAK", nil
	case timeInForce.IsFAS():
		return "FAS", nil
	case timeInForce.IsFOK():
		return "FOK", nil
	case timeInForce.IsSOK():
		return "SOK", nil
	default:
		message := "非対応の執行数量条件です"
		original := fmt.Sprintf("Unexpected time in force code %v", timeInForce.Value())

		myerr := errors.NewMyError(message, original, 500)

		return "", myerr
	}
}
