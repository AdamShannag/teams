package validator

import (
	"context"
	"team-service/validation/violation"
)

type Validator[T any] interface {
	Validate(request T, ctx context.Context) []violation.Violation
}
