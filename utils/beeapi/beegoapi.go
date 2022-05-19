package beeapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/astaxie/beego"
)

type RequestBody struct {
	Ctx          url.Values `json:"ctx"`
	TableName    string     `json:"tablename"`
	Columns      []string   `json:"column"`
	Order        []string   `json:"order"`
	SearchFilter []string   `json:"searchfilter"`
}

type ResponseBeeApi struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type ResponseBeeApiLogin struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    DataLogin `json:"data"`
}
type DataLogin struct {
	Nama  string `json:"nama"`
	Email string `json:"email"`
	Token string `json:"token"`
}
type DataBee struct {
	Uid      string ` form:"uid" json:"uid"`
	Nama     string ` form:"nama" json:"nama"`
	Alamat   string ` form:"alamat" json:"alamat"`
	Email    string ` form:"email" json:"email"`
	Password string ` form:"password" json:"-"`
}
type ResponseBeeApiAjax struct {
	Draw            int32       `json:"draw"`
	RecordsTotal    int64       `json:"recordsTotal"`
	RecordsFiltered int32       `json:"recordsFiltered"`
	Data            interface{} `json:"data"`
}

var client = &http.Client{}

var baseURL = beego.AppConfig.String("baseURL")

func GetAll(token string, reqBody RequestBody) (ResponseBeeApiAjax, error) {
	var data ResponseBeeApiAjax

	json_data, err := json.Marshal(&reqBody)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ajaxgetALL =", token)
	request, err := http.NewRequest("POST", baseURL+"user/ajax", bytes.NewBuffer(json_data))
	if err != nil {
		return ResponseBeeApiAjax{}, err
	}
	request.Header.Set("Content-Type", "application/json")

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))

	res, err2 := client.Do(request)
	if err2 != nil {
		return data, err2
	}

	defer res.Body.Close()

	err3 := json.NewDecoder(res.Body).Decode(&data)
	if err3 != nil {
		return data, err3
	}

	return data, nil

}

func GetById(uid string) (ResponseBeeApi, error) {
	var data ResponseBeeApi

	request, err1 := http.NewRequest("GET", baseURL+"user/"+uid, nil)
	if err1 != nil {
		return data, err1
	}

	response, err2 := client.Do(request)
	if err2 != nil {
		return data, err2
	}
	defer response.Body.Close()

	err3 := json.NewDecoder(response.Body).Decode(&data)
	if err3 != nil {
		return data, err3
	}

	return data, nil
}

func CreateUser(token string, reqBody DataBee) (ResponseBeeApi, error) {
	var data ResponseBeeApi

	json_data, err := json.Marshal(&reqBody)

	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", baseURL+"user", bytes.NewBuffer(json_data))
	if err != nil {
		return ResponseBeeApi{}, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))

	res, err2 := client.Do(request)
	if err2 != nil {
		return data, err2
	}

	defer res.Body.Close()

	err3 := json.NewDecoder(res.Body).Decode(&data)
	if err3 != nil {
		return data, err3
	}

	return data, nil

}

func UpdateUser(token string, reqBody DataBee) (ResponseBeeApi, error) {
	var data ResponseBeeApi

	json_data, err := json.Marshal(&reqBody)

	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("PUT", baseURL+"user/"+reqBody.Uid, bytes.NewBuffer(json_data))
	if err != nil {
		return ResponseBeeApi{}, err
	}
	request.Header.Set("Content-Type", "application/json")

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))

	res, err2 := client.Do(request)
	if err2 != nil {
		return data, err2
	}

	defer res.Body.Close()

	err3 := json.NewDecoder(res.Body).Decode(&data)
	if err3 != nil {
		return data, err3
	}

	return data, nil

}

func DeleteUser(token, uid string) (ResponseBeeApi, error) {
	var data ResponseBeeApi

	request, err := http.NewRequest("DELETE", baseURL+"user/"+uid, nil)
	if err != nil {
		return ResponseBeeApi{}, err
	}
	request.Header.Set("Content-Type", "application/json")

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))

	res, err2 := client.Do(request)
	if err2 != nil {
		return data, err2
	}

	defer res.Body.Close()

	err3 := json.NewDecoder(res.Body).Decode(&data)
	if err3 != nil {
		return data, err3
	}

	return data, nil

}
func UserRegister(reqBody DataBee) (ResponseBeeApi, error) {
	var data ResponseBeeApi

	json_data, err := json.Marshal(&reqBody)

	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", baseURL+"auth/register", bytes.NewBuffer(json_data))
	if err != nil {
		return ResponseBeeApi{}, err
	}
	request.Header.Set("Content-Type", "application/json")

	res, err2 := client.Do(request)
	if err2 != nil {
		return data, err2
	}

	defer res.Body.Close()

	err3 := json.NewDecoder(res.Body).Decode(&data)
	if err3 != nil {
		return data, err3
	}

	return data, nil

}

func UserLogin(reqBody DataBee) (ResponseBeeApiLogin, error) {
	var data ResponseBeeApiLogin

	json_data, err := json.Marshal(&reqBody)

	if err != nil {
		log.Fatal(err)
	}

	request, err := http.NewRequest("POST", baseURL+"auth/login", bytes.NewBuffer(json_data))
	if err != nil {
		return ResponseBeeApiLogin{}, err
	}
	request.Header.Set("Content-Type", "application/json")

	res, err2 := client.Do(request)
	if err2 != nil {
		return data, err2
	}

	defer res.Body.Close()

	err3 := json.NewDecoder(res.Body).Decode(&data)
	if err3 != nil {
		return data, err3
	}

	return data, nil

}
