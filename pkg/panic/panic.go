package httpException

import (
	"errors"
	"fmt"
	"learn-go/internal/constant"
)

func PanicException_(key string, message string) {
	err := errors.New(message)
	err = fmt.Errorf("%s: %w", key, err)
	if err != nil {
		panic(err)
	}
}

func PanicException(responseKey constant.ResponseStatus, message ...string) {
	if condition := len(message) > 0; condition {
		PanicException_(responseKey.GetResponseStatus(), message[0])
		return
	}
	PanicException_(responseKey.GetResponseStatus(), responseKey.GetResponseMessage())
}
