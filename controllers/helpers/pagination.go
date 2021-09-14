package helpers

import (
	"math"
	"strconv"
)

type Pagination struct {
	TotalItems  int `json:"total_items"`
	TotalPages  int `json:"total_pages"`
	CurrentPage int `json:"current_page"`
	PageLimit   int `json:"page_limit"`
}

func Paginate(nItems, page, limit int) Pagination {
	return Pagination{
		TotalItems:  nItems,
		TotalPages:  int(math.Ceil(float64(nItems) / float64(limit))),
		CurrentPage: page,
		PageLimit:   limit,
	}
}

func OffsetLimit(page, limit string) (int, int) {
	var limitInt, offset int
	pageInt, err := strconv.Atoi(page)
	if err == nil {
		pageInt = 1
	}
	limitInt, err = strconv.Atoi(limit)
	if err == nil {
		limitInt = 10
	}
	offset = limitInt * (pageInt - 1)
	return limitInt, offset
}
