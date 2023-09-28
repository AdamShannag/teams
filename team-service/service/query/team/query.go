package team

import (
	"net/http"
	"team-service/filter/team"
	"team-service/service/pagination"
	"team-service/service/sorting"
)

type Query struct {
	Filter     team.Filter
	Sort       sorting.Sort
	Pagination pagination.Pagination
	IsAll      bool
}

func NewQuery(r *http.Request) *Query {
	query := r.URL.Query()
	return &Query{
		IsAll:  isAll(r.Header),
		Filter: *team.NewFilter(query),
		Sort:   *sorting.NewSort(query),
		Pagination: *pagination.NewPagination(
			query.Get("page"),
			query.Get("size"),
		),
	}
}

func isAll(header http.Header) bool {
	return header.Get("all") == "true"
}
