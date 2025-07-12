package api

import (
	"GoTask/task04/forms"
	"GoTask/task04/global"
	"GoTask/task04/model"
	"GoTask/task04/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
	"strings"
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
	//id := context.Param("id")
	idStr := context.Param("id")
	id := utils.ParamToUint(idStr)
	var post model.Post
	global.DB.First(&post, id)
	if post.ID == 0 {
		context.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "文章不存在",
		})
	}
	post = model.Post{
		Title:   utils.ToNullString(postForm.Title),
		Content: utils.ToNullString(postForm.Content),
	}
	//global.DB.Model(&post).Where("id = ? and user_id=?", id, context.MustGet("userId").(uint)).Updates(&post)
	global.DB.Model(&post).Where(&model.Post{Model: gorm.Model{ID: id}, UserId: context.MustGet("userId").(uint)}).Updates(&post)
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "更新文章成功",
	})
}

func DeletePost(context *gin.Context) {
	idStr := context.Param("id")
	id := utils.ParamToUint(idStr)
	tx := global.DB.Begin()
	result := global.DB.Where(&model.Post{Model: gorm.Model{ID: id}, UserId: context.MustGet("userId").(uint)}).Delete(&model.Post{})
	if result.Error != nil {
		tx.Rollback()
		context.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "删除文章失败",
		})
		return
	}
	result = global.DB.Where(&model.Comment{PostId: id}).Delete(&model.Comment{})
	if result.Error != nil {
		tx.Rollback()
		context.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "删除文章失败",
		})
		return
	}
	tx.Commit()
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "删除文章成功",
	})
	return
}

func GetPostDetail(context *gin.Context) {
	idStr := context.Param("id")
	id := utils.ParamToUint(idStr)
	userId := context.MustGet("userId").(uint)
	var post model.Post
	global.DB.Where(&model.Post{Model: gorm.Model{ID: id}, UserId: userId}).Find(&post)
	//post.User = nil
	var user model.User
	global.DB.First(&user, post.UserId)
	post.User = &user
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "获取文章详情成功",
		"data":    post,
	})

}

func GetPostList(context *gin.Context) {
	var postQuery forms.PostQuery
	if err := context.ShouldBindJSON(&postQuery); err != nil {
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
	db := global.DB.Model(&model.Post{})
	if strings.TrimSpace(postQuery.Title) != "" {
		db = db.Where("title like ?", "%"+postQuery.Title+"%")
	}
	if strings.TrimSpace(postQuery.Content) != "" {
		db = db.Where("content like ?", "%"+postQuery.Content+"%")
	}
	offset := (postQuery.PageNumber - 1) * postQuery.PageSize
	limit := postQuery.PageSize
	db = db.Offset(offset).Limit(limit)
	var posts []model.Post
	if err := db.Find(&posts).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "数据库查询失败"})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "获取文章列表成功",
		"data":    posts,
	})
}

func GetCommentList(context *gin.Context) {

}
