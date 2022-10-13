package middleware

import (
	"errors"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, err error) {
	if err != nil {
		var ce *CustomError
		if errors.As(err, &ce) {
			http.Error(w, ce.GetMessage(), ce.GetStatus())
		} else {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
