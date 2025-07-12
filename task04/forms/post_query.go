package forms

type PostQuery struct {
	PageNumber int    `form:"pageNumber" json:"pageNumber" binding:"required,min=1"`
	PageSize   int    `form:"pageSize" json:"pageSize" binding:"required,min=1,max=50"`
	Title      string `form:"title" json:"title"`
	Content    string `form:"content" json:"content"`
}
