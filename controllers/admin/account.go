package admin

type AccountController struct {
	BaseController
}

func (this *AccountController)Get()  {
	this.TplName = "admin/login.html"
}

func (this *AccountController)Post()  {
	this.Data["json"] = "{'result':false,'msg':'密码错误'}"
	this.ServeJSON()
	this.StopRun()
}