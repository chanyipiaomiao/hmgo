package hmgo

// 分页信息
type Page struct {
	PageNo     int  `json:"page_no"`
	PageSize   int  `json:"page_size"`
	TotalPage  int  `json:"total_page"`
	TotalCount int  `json:"total_count"`
	FirstPage  bool `json:"first_page"`
	LastPage   bool `json:"last_page"`
}

// PageUtil生成分页结构工具函数
func PageUtil(count int, pageNo int, pageSize int) Page {

	if pageSize == 0 {
		pageSize = 5
	}

	totalPage := count / pageSize

	if count%pageSize > 0 || count == 0 {
		totalPage = count/pageSize + 1
	}

	return Page{
		PageNo:     pageNo,
		PageSize:   pageSize,
		TotalPage:  totalPage,
		TotalCount: count,
		FirstPage:  pageNo == 1,
		LastPage:   pageNo == totalPage,
	}
}
