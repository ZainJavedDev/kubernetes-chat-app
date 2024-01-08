package controllers

import "github.com/astaxie/beego"

type CheckController struct {
	beego.Controller
}

func (c *CheckController) Post() {
	response := map[string]string{
		"Check": "Passed",
	}
	c.Data["json"] = response
	c.ServeJSON()
}
