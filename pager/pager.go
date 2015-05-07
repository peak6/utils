package pager

import (
	"math"
)

type Pager struct {
	Total     int         `json:"total"`
	CurPage   int         `json:"cur_page"`
	TotalPage int         `json:"total_page"`
	Items     interface{} `json:"items"`
}

//GetPagination return total page and current page
func Page(limit, skip, total int) (int, int) {
	if total == 0 {
		return 0, 0
	}

	var page, totalPage float64
	if skip == 0 {
		skip = 1
	}

	totalPage = math.Ceil((float64(total) / float64(limit)))
	page = float64(skip) * totalPage / float64(total)

	return int(math.Ceil(page)), int(totalPage)
}

func Offset(limit, page int) int {
	if page < 2 {
		return 0
	}

	return ((page - 1) * limit) + 1
}

func New(limit, skip, total int, items interface{}) *Pager {
	curPage, totalPage := Page(limit, skip, total)
	return &Pager{
		Total:     total,
		CurPage:   curPage,
		TotalPage: totalPage,
		Items:     items,
	}
}
