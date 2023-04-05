package errs

import (
	"errors"
	"strings"
)

func Is(err error, msg string) bool {
	return err != nil && strings.Contains(strings.ToLower(err.Error()), strings.ToLower(msg))
}

func Not(err error, msg string) bool {
	return err != nil && !strings.Contains(strings.ToLower(err.Error()), strings.ToLower(msg))
}

func IsErr(err error, target error) bool {
	return err != nil && errors.Is(err, target)
}

func NotErr(err error, target error) bool {
	return err != nil && !errors.Is(err, target)
}
