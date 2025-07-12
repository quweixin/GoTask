package forms

type CommentForm struct {
	Content string `form:"content" json:"content" binding:"required,min=10,max=200"`
	PostId  uint   `form:"post_id" json:"post_id" binding:"required"`
}
