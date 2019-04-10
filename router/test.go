package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../data"
	"../mylog"
	"../tools"
)

var goodstext string = "undefine"

func Test_connect(w http.ResponseWriter, r *http.Request) {
	tools.SetHeader(w)
	tools.WriteJson(w, "connnect scuess!!!")
}

//send back goods message
func TestData(w http.ResponseWriter, r *http.Request) {
	tools.SetHeader(w)
	fmt.Println("1")
	tools.WriteJson(w, data.MockData)
}

//send back type message of goods
func TestData2(w http.ResponseWriter, r *http.Request) {
	tools.SetHeader(w)
	fmt.Println("2")
	tools.WriteJson(w, data.MockData2)
}

//return data of goods detail page
func TestData3(w http.ResponseWriter, r *http.Request) {
	tools.SetHeader(w)
	fmt.Println("3")
	if r.Method != "GET" {
		return
	}
	vars := r.URL.Query()
	tag := vars["tag"][0]
	// id := vars["goodsid"][0]
	switch tag {
	case "base":
		tools.WriteJson(w, data.MockData3)
	case "state":
		tools.WriteJson(w, data.MockData4)
	case "text":
		tools.WriteJson(w, goodstext)
	default:
		return
	}
	return
}

//receive the goods-descriobe-text user upload
func TestData4(w http.ResponseWriter, r *http.Request) {
	tools.SetHeader(w)
	if r.Method != "POST" {
		return
	}
	var postbody map[string]string
	body, err := ioutil.ReadAll(r.Body)
	tools.HandleError("ioutil.readall error", err, 1)
	json.Unmarshal(body, &postbody)
	if postbody["goodsdata"] != "" {
		goodstext = postbody["goodsdata"]
		fmt.Println(goodstext)
		tools.WriteJson(w, "scuess!")
	} else {
		tools.WriteJson(w, "fall")
	}

	return
}

//receive the goods-data from upload-goods-page
func GetGoodsData(w http.ResponseWriter, r *http.Request) {
	tools.SetHeader(w)
	if r.Method != "POST" {
		return
	}
	goodsdata := data.UploadGoodsData{}
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &goodsdata)
	if err != nil {
		mylog.Errorlog.Println(err)
		tools.WriteJson(w, -1)
	}
	fmt.Println(goodsdata)
	tools.WriteJson(w, 1)
}
