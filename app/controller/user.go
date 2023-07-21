package controller

import (
	e2 "go-admin/app/global/e"
	"go-admin/app/global/request"
	"go-admin/app/global/response"
	"go-admin/app/model"
	"go-admin/app/service"
	"go-admin/app/util/validate"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserEdit(c *gin.Context) {
	if c.Request.Method == "GET" {
		pro := service.GetProfile(c)
		if pro.Avatar == "" {
			pro.Avatar = e2.DefaultAvatar
		}
		info, _ := service.GetLoginInfo(c)
		c.HTML(http.StatusOK, "user_show.html", gin.H{"user": pro, "info": info})
	} else {
		var pro request.ProfileForm
		if err := c.ShouldBind(&pro); err != nil {
			response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperEdit).Log(e2.ProfileEdit, c.Request.PostForm).WriteJsonExit()
			return
		}
		if err := service.ProfileEditService(pro, c); err != nil {
			response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperEdit).Log(e2.ProfileEdit, c.Request.PostForm).WriteJsonExit()
			return
		}
		response.SuccessResp(c).SetType(model.OperEdit).Log(e2.ProfileEdit, c.Request.PostForm).WriteJsonExit()
	}
}

func UploadPage(c *gin.Context) {
	c.HTML(http.StatusOK, "upload_profile.html", nil)
}

func AvatarEdit(c *gin.Context) {
	var f request.AvatarForm
	if err := c.ShouldBind(&f); err != nil {
		response.ErrorResp(c).SetMsg(validate.GetValidateError(err)).SetType(model.OperEdit).Log(e2.AvatarEdit, c.Request.PostForm).WriteJsonExit()
		return
	}
	err := service.UpdateAvatarService(f.Avatar, c)
	if err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperEdit).Log(e2.AvatarEdit, c.Request.PostForm).WriteJsonExit()
		return
	}
	response.SuccessResp(c).SetMsg(f.Avatar).SetType(model.OperEdit).Log(e2.AvatarEdit, c.Request.PostForm).WriteJsonExit()
}

func PwdEdit(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "user_pwd.html", gin.H{})
	} else {
		var pwd request.PasswordForm
		if err := c.ShouldBind(&pwd); err != nil {
			response.ErrorResp(c).SetMsg(validate.GetValidateError(err)).SetType(model.OperEdit).Log(e2.PwdEditHandler, c.Request.PostForm).WriteJsonExit()
			return
		}
		if err := service.PwdEditHandlerService(pwd, c); err != nil {
			response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperEdit).Log(e2.PwdEditHandler, c.Request.PostForm).WriteJsonExit()
			return
		}
		response.SuccessResp(c).SetType(model.OperEdit).Log(e2.PwdEditHandler, c.Request.PostForm).WriteJsonExit()
	}
}
