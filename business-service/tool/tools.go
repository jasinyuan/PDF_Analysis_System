package tool

import (
	"businessservice/model"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* 用户账号判断 */
func SearchDb(u model.User, ctx *gin.Context,direction string) {
	result := model.DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&u)
	if direction == "login" {
		login_judge(result,ctx)
	} else {
		register_judge(u,result,ctx)
	}
}

/* 登录账号判断 */
func login_judge(result *gorm.DB,ctx *gin.Context){
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusOK, gin.H{
				"err": "用户名或密码错误",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"查询失败": result.Error,
			})
		}
	} else {
		ctx.Redirect(http.StatusMovedPermanently, "/upload")
	}
}

/* 注册账号判断 */
func register_judge(u model.User,result *gorm.DB,ctx *gin.Context)  {
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			model.DB.Create(&u)
			ctx.Redirect(http.StatusMovedPermanently,"/register_success")
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"数据库插入": result.Error,
			})
		}
	} else {
		ctx.JSON(http.StatusOK,gin.H{
			"err":"该账号已经存在请直接登录",
		})
	}
}