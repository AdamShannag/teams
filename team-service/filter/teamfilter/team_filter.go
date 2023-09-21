package teamfilter

import (
	"entgo.io/ent/dialect/sql"
	"net/url"
	"team-service/api/endpoints/teams"
	"team-service/repository/ent/predicate"
	t "team-service/repository/ent/team"
)

// Filter for search.
type Filter struct {
	Predicate []predicate.Team
}

func NewFilter(query url.Values) *Filter {
	var ps = []predicate.Team{}
	for key := range query {
		if key == teams.TEAM_ID {
			key = t.FieldID
		}
		if t.ValidColumn(key) {
			ps = append(ps, predicate.Team(sql.FieldContains(key, query.Get(key))))
		}
	}
	return &Filter{Predicate: ps}
}
