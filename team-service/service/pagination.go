package service

import (
	"github.com/mohammadyaseen2/pagination/model"
	"github.com/mohammadyaseen2/pagination/pagination"
	"strconv"
)

type Pagination struct {
	Page int
	Size int
}

func NewPagination(pageStr, sizeStr string) *Pagination {
	page := toInt(pageStr)
	size := toInt(sizeStr)

	return &Pagination{
		Size: size,
		Page: page * size,
	}
}

func (p *Pagination) GetResource(currentTotalElement, totalElements int) *model.Resource {
	paginationInfo := pagination.NewPaginationInfo(totalElements, p.createPagination())
	paginationInfo.SetAllData(currentTotalElement)
	return paginationInfo.Resource
}

func (p *Pagination) createPagination() *model.Pagination {
	size := p.Size
	page := 0
	if size != 0 {
		page = p.Page / size
	}

	return pagination.NewPagination(page, size)
}

func toInt(str string) int {
	if str == "" {
		return 0
	}
	i, _ := strconv.Atoi(str)
	return i
}
