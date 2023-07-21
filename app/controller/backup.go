package controller

import (
	"go-admin/app/global/request"
	"go-admin/app/global/response"
	"go-admin/app/service"
	"go-admin/app/util/validate"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TaskList(c *gin.Context) {
	c.HTML(http.StatusOK, "task_list.html", nil)
}

func TaskJson(c *gin.Context) {
	var f request.TaskPage
	if err := c.ShouldBind(&f); err != nil {
		response.ErrorResp(c).SetMsg(validate.GetValidateError(err)).WriteJsonExit()
		return
	}
	data, count, err := service.TaskJson(f)
	if err != nil {
		response.SuccessResp(c).SetCode(0).SetMsg(err.Error()).SetCount(count).WriteJsonExit()
		return
	}
	response.SuccessResp(c).SetCode(0).SetCount(count).SetData(data).WriteJsonExit()
}

func TaskAdd(c *gin.Context) {
	if c.Request.Method == "GET" {
		servers, _, _ := service.ServerList()
		c.HTML(http.StatusOK, "task_add.html", gin.H{"servers": servers})
	} else {
		var f request.TaskForm
		if err := c.ShouldBind(&f); err != nil {
			response.ErrorResp(c).SetMsg(validate.GetValidateError(err)).WriteJsonExit()
			return
		}
		err := service.TaskAdd(f)
		if err != nil {
			response.ErrorResp(c).SetMsg(err.Error()).WriteJsonExit()
			return
		}
		response.SuccessResp(c).SetMsg("任务创建成功，已在后台运行，请稍后查看文件数量!").WriteJsonExit()
		return
	}
}

func TaskEdit(c *gin.Context) {
	c.HTML(http.StatusOK, "task_edit.html", nil)
}

func TaskDel(c *gin.Context) {

}
