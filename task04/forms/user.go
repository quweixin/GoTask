package forms

type RegisterForm struct {
	Username string `form:"username" json:"username" binding:"required,min=5,max=10"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=20"`
}

type LoginForm struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
