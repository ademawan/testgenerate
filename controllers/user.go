package controllers

import (
	"fmt"
	"testgenerate/models"

	"testgenerate/utils/beeapi"

	"github.com/beego/beego/v2/server/web"
)

// UserController operations for User
type UserController struct {
	web.Controller
}
type RspData struct {
	Draw            int32      `json:"draw"`
	RecordsTotal    int32      `json:"recordsTotal"`
	RecordsFiltered int32      `json:"recordsFiltered"`
	Data            []UserResp `json:"data"`
}
type UserResp struct {
	Id     int64  `json:"id"`
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
	Action string `json:"actions"`
}

// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("AjaxGetList", c.AjaxGetList)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Login", c.Login)
	c.Mapping("Register", c.Register)
	c.Mapping("Logout", c.Logout)
}

func (c *UserController) Get() {
	token := c.GetSession("token_login")
	if token == nil {
		c.Redirect("/user/login", 303)
		return
	}

	c.Data["Aboutme"] = "Ade Mawan"

	c.TplName = "dashboard.html"

}

func (c *UserController) AjaxGetList() {
	token := c.GetSession("token_login")

	str := fmt.Sprintf("%v", token)

	ctx, _ := c.Input()
	res, err := beeapi.GetAll(str, beeapi.RequestBody{Ctx: ctx, TableName: "user", Columns: []string{"uid", "nama", "alamat", "email"}, Order: []string{"nama", "alamat", "email"}, SearchFilter: []string{"nama", "alamat"}})
	if err != nil {
		c.Data["json"] = "failed datatables"

	} else {

		c.Data["json"] = res
	}

	c.ServeJSON()

}

func (c *UserController) Post() {
	// token := c.GetSession("token_login")

	var user UserRegisterRequestFormat
	c.BindForm(&user)

	res, err := beeapi.UserRegister(beeapi.DataBee{Nama: user.Nama, Alamat: user.Alamat, Email: user.Email, Password: user.Password})
	if err != nil {
		panic(err.Error())
	} else {
		c.Data["json"] = res
		c.ServeJSON()

	}

}

func (c *UserController) GetOne() {
	uid := c.Ctx.Input.Param(":id")
	if uid <= "" {
		c.Data["json"] = "error parsing"
		c.ServeJSON()
	}

	res, err := beeapi.GetById(uid)
	if err != nil {
		c.Data["json"] = "error get user"
	} else {
		c.Data["json"] = res.Data

	}
	c.ServeJSON()
}

func (c *UserController) GetAll() {
	fmt.Println("hello")
	c.TplName = "testview.html"

}
func (c *UserController) Put() {
	if c.IsAjax() {
		token := c.GetSession("token_login")
		str := fmt.Sprintf("%v", token)

		var user models.User
		c.BindForm(&user)

		// user.Uid = c.Ctx.Input.Param(":id")

		res, err := beeapi.UpdateUser(str, beeapi.DataBee{Uid: user.Uid, Nama: user.Nama, Alamat: user.Alamat})
		if err != nil {
			panic(err.Error())
		} else {

			c.Data["json"] = res
		}
		c.ServeJSON()

	}

}

func (c *UserController) Delete() {
	token := c.GetSession("token_login")
	str := fmt.Sprintf("%v", token)

	uid := c.Ctx.Input.Param(":id")
	fmt.Println(uid)
	if uid == "" {
		c.Data["json"] = "error parsing"
		c.ServeJSON()
	}

	res, err := beeapi.DeleteUser(str, uid)
	if err != nil {
		c.Data["json"] = "error get user"
	} else {
		c.Data["json"] = res.Data

	}
	c.ServeJSON()
}

func (c *UserController) Login() {

	if c.IsAjax() {
		var user UserRegisterRequestFormat
		c.Ctx.BindForm(&user)

		res, err := beeapi.UserLogin(beeapi.DataBee{Email: user.Email, Password: user.Password})
		if err != nil {
			panic(err.Error())
		} else {

			c.SetSession("token_login", res.Data.Token)
			c.Data["Aboutme"] = res.Data.Nama

			// restoken := c.GetSession("token_login")
			// fmt.Println(restoken)
			c.Data["json"] = "res"
			c.ServeJSON()

		}

	}

	c.Layout = "usercontroller/template.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["AuthLayout"] = "usercontroller/login.tpl"
	c.LayoutSections["Scripts"] = "authtest.tpl"
}

func (c *UserController) Register() {

	res := c.GetString("nama")
	fmt.Println(res, "nama")
	if c.IsAjax() {
		var user UserRegisterRequestFormat
		c.Ctx.BindForm(&user)
		fmt.Println(user, "userrrrrr")

		res, err := beeapi.UserRegister(beeapi.DataBee{Nama: user.Nama, Alamat: user.Alamat, Email: user.Email, Password: user.Password})
		fmt.Println(res)
		if err != nil {
			panic(err.Error())
		} else {
			c.Data["json"] = "res"
			c.ServeJSON()

		}

	}
	c.Data["Aboutme"] = "Ade Mawan"

	c.Layout = "usercontroller/template.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["AuthLayout"] = "usercontroller/register.tpl"
	c.LayoutSections["Scripts"] = "authtest.tpl"

}
func (c *UserController) Logout() {
	c.DestroySession()

	c.Redirect("/user/login", 303)

}
