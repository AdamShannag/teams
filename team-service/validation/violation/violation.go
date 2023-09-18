package violation

type Violation struct {
	Field   string
	Message string
}

func FieldViolation(field string, err error) Violation {
	return Violation{
		Field:   field,
		Message: err.Error(),
	}
}
