package api

import (
	"GoTask/task04/forms"
	"GoTask/task04/global"
	"GoTask/task04/model"
	"GoTask/task04/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func CreatePost(context *gin.Context) {
	var postForm forms.PostForm
	if err := context.ShouldBind(&postForm); err != nil {
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if !ok {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": utils.ErrorFormat(errs.Translate(global.Trans)),
		})
		return
	}
	var post model.Post
	post = model.Post{
		Title:   utils.ToNullString(postForm.Title),
		Content: utils.ToNullString(postForm.Content),
		UserId:  context.MustGet("userId").(uint),
	}
	global.DB.Create(&post)
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "创建文章成功",
	})
}

func UpdatePost(context *gin.Context) {

}

func DeletePost(context *gin.Context) {

}

func GetPostDetail(context *gin.Context) {

}

func GetPostList(context *gin.Context) {

}

func GetCommentList(context *gin.Context) {

}
