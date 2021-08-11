package controllers

import (
	"encoding/json"
	model "go_connectdb/models"
	"net/http"
)

type ResponeUser struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
	Data   []model.Users
}

func GetUser(rw http.ResponseWriter, r *http.Request) {

	//  get
	user_data, err := model.UserData()
	if err != nil {
		// response
		var rs = ResponeUser{
			Status: "err",
			Msg:    err.Error(),
		}
		var user_data, _ = json.Marshal(rs)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(user_data))
		return
	}

	// response
	var rs = ResponeUser{
		Status: "success",
		Msg:    "OK",
		Data:   user_data,
	}
	rs_user_data, _ := json.Marshal(rs)
	rw.WriteHeader(200)
	rw.Write(rs_user_data)

}

func getUser(rw http.ResponseWriter, r *http.Request) {

	//  get
	user_data, err := model.UserData()
	if err != nil {
		// response
		var rs = ResponeUser{
			Status: "err",
			Msg:    err.Error(),
		}
		var user_data, _ = json.Marshal(rs)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(user_data))
		return
	}

	// response
	var rs = ResponeUser{
		Status: "success",
		Msg:    "OK",
		Data:   user_data,
	}
	rs_user_data, _ := json.Marshal(rs)
	rw.WriteHeader(200)
	rw.Write(rs_user_data)

}
