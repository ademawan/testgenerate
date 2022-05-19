package main

import (
	"fmt"
	_ "testgenerate/routers"

	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	mysqlDB := beego.AppConfig.String("mysqlDB")
	mysqlUser := beego.AppConfig.String("mysqlUser")
	mysqlPass := beego.AppConfig.String("mysqlPass")
	conn := fmt.Sprintf("%s:%s@/%s?charset=utf8", mysqlUser, mysqlPass, mysqlDB)

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase(`default`, "mysql", conn)
}

func main() {
	web.SetStaticPath("/images", "static/images")
	web.SetStaticPath("/css", "static/css")
	web.SetStaticPath("/js", "static/js")
	// web.BConfig.WebConfig.Session.SessionOn = true

	web.Run()
}
