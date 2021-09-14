package helpers

import "math"

type Pagination struct {
	TotalItems  int `json:"total_items"`
	TotalPages  int `json:"total_pages"`
	CurrentPage int `json:"current_page"`
	PageLimit   int `json:"page_limit"`
}

func Paginate(nItems, page, limit int) Pagination {
	return Pagination{
		TotalItems: nItems,
		TotalPages: int(math.Ceil(float64(nItems)/float64(limit))),
		CurrentPage: page,
		PageLimit: limit,
	}
}

func Offset(page, limit int) int {
	return limit * (page-1)
}