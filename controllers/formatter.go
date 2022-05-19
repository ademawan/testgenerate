package controllers

import "html/template"

type UserResponse struct {
	Id      int64
	Nama    string
	Alamat  string
	Actions *template.Template
}

type UserRegisterRequestFormat struct {
	Nama     string `form:"nama" json:"nama"`
	Alamat   string `form:"alamat" json:"alamat"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}
