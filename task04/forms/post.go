package forms

type PostForm struct {
	Title   string `form:"title" json:"title" binding:"required,min=5,max=100"`
	Content string `form:"content" json:"content" binding:"required,min=10,max=200"`
}
