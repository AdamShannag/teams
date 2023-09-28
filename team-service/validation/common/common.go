package common

import (
	"errors"
)

func IsNilString(value *string) error {
	if value == nil {
		return errors.New("is required")
	}
	return nil
}

func IsEmptyString(value string) error {
	if len(value) == 0 {
		return errors.New("is required")
	}
	return nil
}

func IsEmpty(value []string) error {
	if value == nil || len(value) == 0 {
		return errors.New("is required")
	}
	return nil
}
