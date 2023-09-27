package sorting

import (
	"entgo.io/ent/dialect/sql"
	"net/url"
	"strings"
	"team-service/constant"
	stringshelper "team-service/helper/strings"
	"team-service/repository/ent"
)

type Sort struct {
	Order func(*sql.Selector)
}

func NewSort(query url.Values) *Sort {
	sortBy := query.Get("sortBy")
	if sortBy == "" || sortBy == constant.TEAM_ID {
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
