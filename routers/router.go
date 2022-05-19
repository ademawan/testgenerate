package routers

import (
	"testgenerate/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.Router("/", &controllers.UserController{}, "get:Get")
	web.Router("/ajax", &controllers.UserController{}, "post:AjaxGetList")

	web.Router("/user/add", &controllers.UserController{}, "post:Post")
	web.Router("/user/save", &controllers.UserController{}, "put:Put")
	web.Router("/user/:id", &controllers.UserController{}, "delete:Delete")

	web.Router("/user/login", &controllers.UserController{}, "get,post:Login")
	web.Router("/user/register", &controllers.UserController{}, "get:Register;post:Register")
	web.Router("/user/logout", &controllers.UserController{}, "get:Logout")

}
