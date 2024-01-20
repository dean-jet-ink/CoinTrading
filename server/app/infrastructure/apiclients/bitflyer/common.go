package bitflyer

import (
	"cointrading/app/domain/entities"
	"cointrading/app/domain/myerror"
	"cointrading/app/domain/valueobjects"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
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

func (b *bitflyerClient) convertSymbolToString(symbol *valueobjects.Symbol) (string, error) {
	switch {
	case symbol.IsBTCJPY():
		return "BTC_JPY", nil
	case symbol.IsETHJPY():
		return "ETH_JPY", nil
	case symbol.IsXRPJPY():
		return "XRP_JPY", nil
	default:
		err := fmt.Errorf("%w: Unexpected symbol code %v", myerror.ErrUnexpectedSymbol, symbol.Value())

		return "", err
	}
}

func (b *bitflyerClient) convertSymbolToInt(symbol string) (int, error) {
	symbolInts := map[string]int{
		"BTC_JPY": valueobjects.BTCJPY.Value(),
		"ETH_JPY": valueobjects.ETHJPY.Value(),
		"XRP_JPY": valueobjects.XRPJPY.Value(),
	}

	symbolInt, ok := symbolInts[symbol]
	if !ok {
		err := fmt.Errorf("%w: Unexpected symbol %v", myerror.ErrUnexpectedSymbol, symbol)

		return 0, err
	}

	return symbolInt, nil
}

func (b *bitflyerClient) convertCurrencyCodeToInt(currencyCode string) (int, bool) {
	currencyCodes := map[string]int{
		"JPY": valueobjects.JPY.Value(),
		"BTC": valueobjects.BTC.Value(),
		"ETH": valueobjects.ETH.Value(),
		"XRP": valueobjects.XRP.Value(),
	}

	code, ok := currencyCodes[currencyCode]
	if !ok {
		return 0, false
	}

	return code, true
}

func (b *bitflyerClient) convertOrderTypeToString(orderType *valueobjects.OrderType) (string, error) {
	switch {
	case orderType.IsMarket():
		return "MARKET", nil
	case orderType.IsLimit():
		return "LIMIT", nil
	case orderType.IsStop():
		return "STOP", nil
	default:
		err := fmt.Errorf("%w: Unexpected order type code: %v", myerror.ErrUnexpectedOrderType, orderType.Value())

		return "", err
	}
}

func (b *bitflyerClient) convertOrderTypeToInt(orderType string) (int, error) {
	orderTypes := map[string]int{
		"MARKET": valueobjects.Market.Value(),
		"LIMIT":  valueobjects.Limit.Value(),
	}

	orderTypeInt, ok := orderTypes[orderType]
	if !ok {
		err := fmt.Errorf("%w: Unexpected order type : %v", myerror.ErrUnexpectedOrderType, orderType)

		return 0, err
	}

	return orderTypeInt, nil
}

func (b *bitflyerClient) convertTimeInForceToString(timeInForce *valueobjects.TimeInForce) (string, error) {
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
		err := fmt.Errorf("%w: Unexpected time in force code %v", myerror.ErrUnexpectedTimeInForce, timeInForce.Value())

		return "", err
	}
}

func (b *bitflyerClient) convertTimeInForceToInt(timeInForce string) (int, error) {
	timeInForces := map[string]int{
		"GTC": valueobjects.GTC.Value(),
		"IOK": valueobjects.IOK.Value(),
		"FOK": valueobjects.FOK.Value(),
	}

	timeInForceInt, ok := timeInForces[timeInForce]
	if !ok {
		err := fmt.Errorf("%w: Unexpected time in force : %v", myerror.ErrUnexpectedTimeInForce, timeInForce)

		return 0, err
	}

	return timeInForceInt, nil
}

func (b *bitflyerClient) convertOrderStatusToInt(status string) (int, error) {
	orderStatuses := map[string]int{
		"ACTIVE":    valueobjects.Waiting.Value(),
		"COMPLETED": valueobjects.Completed.Value(),
		"CANCELED":  valueobjects.Canceled.Value(),
		"EXPIRED":   valueobjects.Expired.Value(),
		"REJECTED":  valueobjects.Rejected.Value(),
	}

	orderStatusInt, ok := orderStatuses[status]
	if !ok {
		err := fmt.Errorf("%w: Unexpected order status : %v", myerror.ErrUnexpectedOrderStatus, status)

		return 0, err
	}

	return orderStatusInt, nil
}

func (b *bitflyerClient) entityToOrder(entity *entities.Order) (*order, error) {
	symbolStr, err := b.convertSymbolToString(entity.Symbol())
	if err != nil {
		return nil, err
	}

	orderTypeStr, err := b.convertOrderTypeToString(entity.OrderType())
	if err != nil {
		return nil, err
	}

	timeInForceStr, err := b.convertTimeInForceToString(entity.TimeInForce())
	if err != nil {
		return nil, err
	}

	return &order{
		ProductCode:    symbolStr,
		ChildOrderType: orderTypeStr,
		Side:           strings.ToUpper(entity.Side().Value()),
		Price:          entity.Price(),
		Size:           entity.Size(),
		MinuteToExpire: 1440,
		TimeInForce:    timeInForceStr,
	}, nil
}

func (b *bitflyerClient) orderToEntity(order *order) (*entities.Order, error) {
	symbol, err := b.convertSymbolToInt(order.ProductCode)
	if err != nil {
		return nil, err
	}

	orderType, err := b.convertOrderTypeToInt(order.ChildOrderType)
	if err != nil {
		return nil, err
	}

	timeInForce, err := b.convertTimeInForceToInt(order.TimeInForce)
	if err != nil {
		return nil, err
	}

	status, err := b.convertOrderStatusToInt(order.Status)
	if err != nil {
		return nil, err
	}

	return entities.NewOrder(
		order.ChildOrderAcceptanceID,
		symbol,
		order.Side,
		orderType,
		order.Price,
		order.Size,
		timeInForce,
		status,
	)
}
