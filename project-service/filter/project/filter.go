package project

import (
	"entgo.io/ent/dialect/sql"
	"net/url"
	"project-service/constant"
	stringshelper "project-service/helper/strings"
	"project-service/repository/ent/predicate"
	"project-service/repository/ent/project"
)

// Filter for search.
type Filter struct {
	Predicate []predicate.Project
}

func NewFilter(query url.Values) *Filter {
	var ps = []predicate.Project{}
	for key := range query {
		if key == constant.PROJECT_ID {
			key = project.FieldID
		}
		key = stringshelper.ToSnakeCase(key)
		if project.ValidColumn(key) {
			ps = append(ps, sql.FieldContains(key, query.Get(key)))
		}
	}
	return &Filter{Predicate: ps}
}
