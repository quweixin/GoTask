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

func CreateComment(context *gin.Context) {

	var commentForm forms.CommentForm
	if err := context.ShouldBindJSON(&commentForm); err != nil {
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

	userId := context.MustGet("userId").(uint)
	comment := model.Comment{
		Content: utils.ToNullString(commentForm.Content),
		PostId:  commentForm.PostId,
		UserId:  userId,
	}
	global.DB.Create(&comment)
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "创建评论成功",
	})

}
