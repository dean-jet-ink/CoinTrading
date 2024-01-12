package middleware

import (
	"cointrading/app/presentation/router"
	"errors"
	"log"
	"net/http"
)

func Recover(c router.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)

			err := errors.New("予期しないエラーが発生しました")

			c.JSON(http.StatusInternalServerError, err)
		}
	}()

	c.Next()
}
