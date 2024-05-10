package errhandler

import (
	"errors"
	"fmt"
)

func Wrap(message string, err error) error {
	return fmt.Errorf("%s: %w", message, err)
}

var (
	ErrNotFound = errors.New("not found")
)
