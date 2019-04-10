package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../mylog"
)

func WriteJson(w http.ResponseWriter, data interface{}) {
	jsondata, err := json.Marshal(data)
	HandleError("worng at tool.go writejson :", err, 1)
	w.Write(jsondata)
}

func SetHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
}

//the simple function to handle the err,use avariable to contral
//exist or containue when error happend
func HandleError(tag string, err error, method int) {
	if err != nil {
		switch method {
		case 1:
			mylog.Errorlog.Println(tag, err)
		case -1:
			mylog.Errorlog.Fatal(tag, err)
		}
	}
}

//equal to  a !="" ? a:b
func NotNullVal(a string, b string) string {
	if a == "" {
		return b
	}
	return a
}

func ShowPostData(r *http.Request) {
	var postbody map[string]interface{}
	body, err := ioutil.ReadAll(r.Body)
	HandleError("showPostData readall() :", err, -1)
	json.Unmarshal(body, &postbody)
	for key, val := range postbody {
		fmt.Println(key, "  -->  ", val)
	}
}
