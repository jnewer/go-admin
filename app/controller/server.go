package controller

import (
	"go-admin/app/global/request"
	"go-admin/app/global/response"
	"go-admin/app/service"
	"go-admin/app/util/gconv"
	"go-admin/app/util/validate"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServerList(c *gin.Context) {
	c.HTML(http.StatusOK, "server_list.html", nil)
}

func ServerJson(c *gin.Context) {
	var f request.TaskServerPage
	if err := c.ShouldBind(&f); err != nil {
		response.ErrorResp(c).SetMsg(validate.GetValidateError(err)).WriteJsonExit()
		return
	}
	data, count, err := service.ServerJson(f)
	if err != nil {
		response.SuccessResp(c).SetCode(0).SetMsg(err.Error()).SetCount(count).WriteJsonExit()
		return
	}
	response.SuccessResp(c).SetCode(0).SetCount(count).SetData(data).WriteJsonExit()
}

func ServerAdd(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "server_add.html", nil)
	} else {
		var f request.TaskServerForm
		if err := c.ShouldBind(&f); err != nil {
			response.ErrorResp(c).SetMsg(validate.GetValidateError(err)).WriteJsonExit()
			return
		}
		err := service.ServerAdd(f)
		if err != nil {
			response.ErrorResp(c).SetMsg(err.Error()).WriteJsonExit()
			return
		}
		response.SuccessResp(c).WriteJsonExit()
		return
	}
}

func ServerEdit(c *gin.Context) {
	if c.Request.Method == "GET" {
		id := c.Query("id")
		s, _ := service.FindServerById(gconv.Int(id))
		c.HTML(http.StatusOK, "server_edit.html", gin.H{"s": s})
	} else {
		var f request.TaskServerForm
		if err := c.ShouldBind(&f); err != nil {
			response.ErrorResp(c).SetMsg(validate.GetValidateError(err)).WriteJsonExit()
			return
		}
		err := service.ServerEdit(f)
		if err != nil {
			response.ErrorResp(c).SetMsg(err.Error()).WriteJsonExit()
			return
		}
		response.SuccessResp(c).WriteJsonExit()
		return
	}
}

func ServerDel(c *gin.Context) {
	id := c.PostForm("id")
	err := service.ServerDel(gconv.Int(id))
	if err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).WriteJsonExit()
		return
	}
	response.SuccessResp(c).WriteJsonExit()
	return
}
