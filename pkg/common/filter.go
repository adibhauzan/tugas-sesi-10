package common

type PaginationFilter struct {
	Page     int
	PageSize int
	Search   string
	SortBy   string
	SortDir  string
}
type PaginationMeta struct {
	Total         int64  `json:"total"`
	Page          int    `json:"page"`
	PageSize      int    `json:"page_size"`
	SortBy        string `json:"sort_by"`
	SortDirection string `json:"sort_direction"`
	Search        string `json:"search"`
}

func (f *PaginationFilter) Normalize() {
	if f.Page <= 0 {
		f.Page = 1
	}
	if f.PageSize <= 0 {
		f.PageSize = 10
	}
	if f.SortDir == "" {
		f.SortDir = "desc"
	}
}
