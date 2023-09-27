package team

import (
	"entgo.io/ent/dialect/sql"
	"net/url"
	"team-service/constant"
	"team-service/helper/strings"
	"team-service/repository/ent/predicate"
	"team-service/repository/ent/team"
)

// Filter for search.
type Filter struct {
	Predicate []predicate.Team
}

func NewFilter(query url.Values) *Filter {
	var ps = []predicate.Team{}
	for key := range query {
		if key == constant.TEAM_ID {
			key = team.FieldID
		}
		key = stringshelper.ToSnakeCase(key)
		if team.ValidColumn(key) {
			ps = append(ps, sql.FieldContains(key, query.Get(key)))
		}
	}
	return &Filter{Predicate: ps}
}
