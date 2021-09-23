package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pear-admin-go/app/global/request"
	"pear-admin-go/app/global/response"
	"pear-admin-go/app/model"
	"pear-admin-go/app/service"
	"pear-admin-go/app/util/e"
	"pear-admin-go/app/util/validate"
)

func UserShow(c *gin.Context) {
	pro := service.GetProfile(c)
	if pro.Avatar == "" {
		pro.Avatar = e.DefaultAvatar
	}
	info, _ := service.GetLoginInfo()
	c.HTML(http.StatusOK, "user_show.html", gin.H{"user": pro, "info": info})
}

func UploadPage(c *gin.Context) {
	c.HTML(http.StatusOK, "upload_profile.html", nil)
}

func AvatarEdit(c *gin.Context) {
	var f request.AvatarForm
	if err := c.ShouldBind(&f); err != nil {
		response.ErrorResp(c).SetMsg("上传失败"+validate.GetValidateError(err)).SetType(model.OperEdit).Log(e.AvatarEdit, c.Request.PostForm).WriteJsonExit()
		return
	}
	err := service.UpdateAvatarService(f.Avatar, c)
	if err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperEdit).Log(e.AvatarEdit, c.Request.PostForm).WriteJsonExit()
		return
	}
	response.SuccessResp(c).SetMsg(f.Avatar).SetType(model.OperEdit).Log(e.AvatarEdit, c.Request.PostForm).WriteJsonExit()
}

func ProfileEdit(c *gin.Context) {
	var pro request.ProfileForm
	if err := c.ShouldBind(&pro); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperEdit).Log(e.ProfileEdit, c.Request.PostForm).WriteJsonExit()
		return
	}
	if err := service.ProfileEditService(pro, c); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperEdit).Log(e.ProfileEdit, c.Request.PostForm).WriteJsonExit()
		return
	}
	response.SuccessResp(c).SetType(model.OperEdit).Log(e.ProfileEdit, c.Request.PostForm).WriteJsonExit()
}

func PwdEdit(c *gin.Context) {
	c.HTML(http.StatusOK, "user_pwd.html", gin.H{})
}

func PwdEditHandler(c *gin.Context) {
	var pwd request.PasswordForm
	if err := c.ShouldBind(&pwd); err != nil {
		response.ErrorResp(c).SetMsg(validate.GetValidateError(err)).SetType(model.OperEdit).Log(e.PwdEditHandler, c.Request.PostForm).WriteJsonExit()
		return
	}
	if err := service.PwdEditHandlerService(pwd, c); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).SetType(model.OperEdit).Log(e.PwdEditHandler, c.Request.PostForm).WriteJsonExit()
		return
	}
	response.SuccessResp(c).SetType(model.OperEdit).Log(e.PwdEditHandler, c.Request.PostForm).WriteJsonExit()
}
