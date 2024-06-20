package util

import (
	"github.com/kataras/iris/v12"
	"reflect"
)

// PaginationData ...
type PaginationData struct {
	Pagination Pagination  `json:"pagination"`
	Objects    interface{} `json:"objects"`
}

// Pagination ...
type Pagination struct {
	// NOTE: 等 C 端列表类旧接口协议更新后，再将 Total 的 filed 加上 omitempty tag
	// NOTE: 目前 C 端前端会根据 total 做判断，故先不做改动
	Total int `json:"total"`
	From  int `json:"from"`
	To    int `json:"to"`
}

// NewPagination ...
func NewPagination(total, from, to int) Pagination {
	return Pagination{total, from, to}
}

// NewPaginationData ...
func NewPaginationData(objects interface{}, pagination Pagination) *PaginationData {
	if objects == nil || reflect.ValueOf(objects).IsNil() {
		// NOTE: convert ‘objects’ of http response body null to []
		objects = []interface{}{}
	}
	return &PaginationData{pagination, objects}
}

// NewPaginationDataWithoutPaged ...
func NewPaginationDataWithoutPaged(objects interface{}) *PaginationData {
	return NewPaginationData(objects, Pagination{})
}

// GetPageData ...
func GetPageData(ctx iris.Context, isQueryAll bool, maxLimit ...int) (limit, offset int) {
	offset = ctx.URLParamIntDefault("offset", 0)
	limit = ctx.URLParamIntDefault("limit", 0)

	// maxLimit 为第一优先级
	if len(maxLimit) > 0 && maxLimit[0] > 0 {
		max := maxLimit[0]
		if limit > max || limit == 0 {
			limit = max
		}
	} else {
		// NOTE: 查全量时 limit 可以为0, 否则默认值为10
		// NOTE: 未设置maxLimit时，限制limit最大值为10
		if !isQueryAll {
			if limit <= 0 || limit > 10 {
				limit = 10
			}
		}
	}
	return
}

// NewPageData ...
func NewPageData(objects interface{}, total int, limit int, offset int) *PaginationData {
	var from, to int
	if limit > 0 && offset >= 0 {
		from = offset
		to = int(offset + limit)
	} else {
		from = 0
		if total > 0 {
			to = total - 1
		}
	}
	return NewPaginationData(objects, NewPagination(total, from, to))
}

// NewPageDataWithoutTotal ...
func NewPageDataWithoutTotal(objects interface{}, limit int, offset int) *PaginationData {
	var from, to int
	if limit > 0 && offset >= 0 {
		from = offset
		to = offset + limit
	} else {
		from = 0
		switch reflect.TypeOf(objects).Kind() {
		case reflect.Slice, reflect.Array:
			if total := reflect.ValueOf(objects).Len(); total > 0 {
				to = total - 1
			}
		}
	}
	return NewPaginationData(objects, NewPagination(0, from, to))
}
