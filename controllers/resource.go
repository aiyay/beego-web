package controllers

import (
	"beegoWeb/service"
	"beegoWeb/utils"
	"strconv"

	"github.com/astaxie/beego"
)

type ResourceController struct {
	beego.Controller
}

func (c *ResourceController) Add() {
	c.Data["Entitys"] = service.GetAllResources()
	c.TplName = "resource/add.html"
}

func (c *ResourceController) Adddo() {
	parentId, _ := strconv.Atoi(c.GetString("parentId"))
	name := c.GetString("name")
	controller := c.GetString("controller")
	action := c.GetString("action")
	sort, _ := strconv.Atoi(c.GetString("sort"))
	url := c.GetString("url")
	isShow, _ := strconv.Atoi(c.GetString("isShow"))
	isExist := service.IsExistResource(controller, action)
	if isExist == true {
		beego.Info("资源已存在")
	} else {
		service.AddResource(parentId, name, controller, action, sort, url, isShow)
	}
	c.Data["Entitys"] = service.GetAllResources()
	c.TplName = "resource/add.html"
	c.Redirect("/resource/index", 301)
}

func (c *ResourceController) Edit() {
	c.TplName = "resource/edit.html"
}

func (c *ResourceController) Editdo() {
	c.TplName = "resource/edit.html"
}

func (c *ResourceController) Delete() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	utils.CheckError(err)
	service.DeleteResrouce(id)
	c.TplName = "resource/index.html"
	c.Redirect("/resource/index", 301)
}

func (c *ResourceController) Index() {
	pageIndex := 0
	ci := c.Input().Get("pageIndex")
	if ci != "" {
		pi, err := strconv.Atoi(ci)
		utils.CheckError(err)
		pageIndex = pi
	}
	pageSize := 10
	users, totalCount := service.GetResources(pageIndex, pageSize)
	c.Data["PageIndex"] = pageIndex
	c.Data["PageSize"] = pageSize
	c.Data["TotalCount"] = totalCount
	c.Data["Entitys"] = users
	c.TplName = "resource/index.html"
}
