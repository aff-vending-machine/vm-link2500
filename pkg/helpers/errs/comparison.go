package errs

import (
	"errors"
	"strings"
)

func HasMsg(err error, msg string) bool {
	if err == nil {
		return false
	}

	return strings.Contains(strings.ToLower(err.Error()), strings.ToLower(msg))
}

func NoMsg(err error, msg string) bool {
	if err == nil {
		return true
	}

	return !strings.Contains(strings.ToLower(err.Error()), strings.ToLower(msg))
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func Not(err, target error) bool {
	return !errors.Is(err, target)
}
