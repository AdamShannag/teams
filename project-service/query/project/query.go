package project

import (
	"net/http"
	"project-service/filter/project"
	"project-service/service/pagination"
	"project-service/service/sorting"
)

type Query struct {
	Filter     project.Filter
	Sort       sorting.Sort
	Pagination pagination.Pagination
	IsAll      bool
}

func NewQuery(r *http.Request) *Query {
	query := r.URL.Query()
	return &Query{
		IsAll:  isAll(r.Header),
		Filter: *project.NewFilter(query),
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
