package controllers

import (
	"github.com/astaxie/beego"
//    "htz_blog/models"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Get() {
	c.TplNames = "admin.html"
}

func (c *AdminController) Login() {
	c.TplNames = "admin.html"
}

func (c *AdminController) Post() {
//	title := c.GetString("title")
//	content := c.GetString("content")
//	kind := c.GetString("kind")
//	ctime := c.GetString("ctime")
//	err := models.ArticleModel().Add(title, content, kind, ctime)
//	if nil != err {
//		c.Abort("500")
//	}

	c.Data["kinds"] = []string{"经典", "学生问答", "心得体会"}
	c.TplNames = "admin.html"
}


