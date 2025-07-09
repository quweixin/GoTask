package api

import (
	"GoTask/task04/forms"
	"GoTask/task04/global"
	"GoTask/task04/middlewares"
	"GoTask/task04/model"
	"GoTask/task04/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func Register(context *gin.Context) {
	var registerForm forms.RegisterForm
	if err := context.ShouldBind(&registerForm); err != nil {
		//errs,ok :=err.(validator.ValidationErrors)
		// 上面写法 等于下面用法
		var errs validator.ValidationErrors
		// 类型转换 如果是 validator.ValidationErrors 类型的错误，则进行翻译
		ok := errors.As(err, &errs)
		if !ok {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": utils.ErrorFormat(errs.Translate(global.Trans)),
		})
		return
	}
	var user model.User
	//global.DB.Model(&user).Where("user_name = ? or email = ?", registerForm.Username, registerForm.Email).First(&user)
	global.DB.Where(&model.User{Email: registerForm.Email}).Or(&model.User{Username: registerForm.Username}).First(&user)
	if user.ID != 0 {
		if user.Username == registerForm.Username {
			context.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"success": false,
				"message": "用户名已存在",
			})
			return
		}
		if user.Email == registerForm.Email {
			context.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"success": false,
				"message": "email已注册",
			})
			return
		}
	}
	user = model.User{
		Email:    registerForm.Email,
		Password: utils.HashUtil(registerForm.Password),
		Username: registerForm.Username,
	}
	global.DB.Create(&user)

	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"success": true,
		"message": "注册成功",
	})
}

func Login(context *gin.Context) {
	var loginForm forms.LoginForm
	if err := context.ShouldBind(&loginForm); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user model.User

	global.DB.Where(&model.User{Username: loginForm.Username}).First(&user)
	if user.ID == 0 {
		context.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": false,
			"message": "用户不存在",
		})
		return
	}

	if !utils.CheckPasswordHash(loginForm.Password, user.Password) {
		context.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": false,
			"message": "密码错误",
		})
		return
	}

	token, err := middlewares.GenerateToken(user.ID, user.Username)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "生成token失败",
			"error":   err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "登录成功",
		"data":    gin.H{"token": token},
	})
}
