package controllers

import (
	"github.com/astaxie/beego"
)

type ActivityController struct {
	beego.Controller
}

func (c *ActivityController) TurnaroundDraw() {
	c.TplName = "activity/turnaround-draw.html"
}

func (c *ActivityController) TurnaroundDrawDo() {
	c.TplName = "activity/turnaround-draw.html"
}
