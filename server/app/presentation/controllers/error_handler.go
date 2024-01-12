package controllers

import (
	"cointrading/app/domain/myerror"
	"cointrading/app/presentation/router"
	"errors"
	"log"
	"net/http"
)

type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}

func handleError(err error, c router.Context) {
	log.Println(err)

	e := errors.Unwrap(err)

	if e != nil {
		e = err
	}

	resp := &ErrorResponse{
		ErrorMessage: e.Error(),
	}

	if errors.Is(myerror.ErrNotFoundOrder, e) ||
		errors.Is(myerror.ErrBadRequest, e) {
		// 400
		c.JSON(http.StatusBadRequest, resp)
	} else if errors.Is(myerror.ErrInvalidOrderSide, e) ||
		errors.Is(myerror.ErrUnexpectedCurrencyCode, e) ||
		errors.Is(myerror.ErrUnexpectedExchange, e) ||
		errors.Is(myerror.ErrUnexpectedOrderStatus, e) ||
		errors.Is(myerror.ErrUnexpectedOrderType, e) ||
		errors.Is(myerror.ErrUnexpectedSymbol, e) ||
		errors.Is(myerror.ErrUnexpectedTimeInForce, e) ||
		errors.Is(myerror.ErrFailedToConnectNetwork, e) {
		// 500
		c.JSON(http.StatusInternalServerError, e)
	} else {
		// 500
		resp = &ErrorResponse{
			ErrorMessage: "予期しないエラーが発生しました",
		}

		c.JSON(http.StatusInternalServerError, resp)
	}
}
