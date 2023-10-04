package sorting

import (
	"entgo.io/ent/dialect/sql"
	"net/url"
	"project-service/constant"
	stringshelper "project-service/helper/strings"
	"project-service/repository/ent"
	"strings"
)

type Sort struct {
	Order func(*sql.Selector)
}

func NewSort(query url.Values) *Sort {
	sortBy := query.Get("sortBy")
	if sortBy == "" || sortBy == constant.PROJECT_ID {
		sortBy = "id"
	}
	sortOrder := query.Get("sortOrder")
	order := getSortOrder(sortOrder)(stringshelper.ToSnakeCase(sortBy))
	return &Sort{
		Order: order,
	}
}

func getSortOrder(order string) func(fields ...string) func(*sql.Selector) {
	switch strings.ToLower(order) {
	case "desc":
		return ent.Desc
	default:
		return ent.Asc
	}
}
