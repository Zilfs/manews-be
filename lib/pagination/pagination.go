package pagination

import (
	"manews/internal/core/domain/entity"
	"math"
)

type PaginationInterface interface {
	AddPagination(totalData, page, perPage int) (*entity.Page, error)
}

type Options struct{}

func (o *Options) AddPagination(totalData, page, perPage int) (*entity.Page, error) {
	newPage := page

	if newPage <= 0 {
		return nil, ErrorPage
	}

	linmitData := 10
	if perPage > 0 {
		linmitData = perPage
	}

	totalPage := int(math.Ceil(float64(totalData) / float64(linmitData)))

	last := (linmitData * newPage)
	first := last - linmitData

	if totalData < last {
		last = totalData
	}

	zeroPage := &entity.Page{PageCount: 1, Page: newPage}
	if totalPage == 0 && newPage == 1 {
		return zeroPage, nil
	}

	if newPage > totalPage {
		return nil, ErrorMaxPage
	}

	pages := &entity.Page{
		Page:       newPage,
		Perpage:    perPage,
		PageCount:  totalPage,
		TotalCount: totalData,
		First:      first,
		Last:       last,
	}

	return pages, nil
}

func NewPagination() PaginationInterface {
	pagination := new(Options)

	return pagination
}
