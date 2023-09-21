package common

import "errors"

func IsEmptyString(value string) error {
	if len(value) == 0 {
		return errors.New("is required")
	}
	return nil
}
