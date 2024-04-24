package dto

// pagination dto
type PaginationDto struct {
	Page     *int `json:"page" form:"page"`
	PageSize *int `json:"pageSize" form:"pageSize"`
}

// Init initializes the PaginationDto, setting fields to nil if they are not set
func (p *PaginationDto) Init() {
	if p.Page != nil && *p.Page == 0 {
		p.Page = nil
	}
	if p.PageSize != nil && *p.PageSize == 0 {
		p.PageSize = nil
	}
}
