package router

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	"../config"
	"../tools"
)

var imgpath string = "./source/images/" //the directory path of save and get images

//receive the cover-images of a goods and return the link of see it
func GetCover(w http.ResponseWriter, r *http.Request) {
	tools.SetHeader(w)
	if r.Method != "POST" {
		return
	}
	//parse from and get the base message
	var username = r.PostFormValue("name")
	r.ParseMultipartForm(1002 << 10)
	//check if file is empty
	if len(r.MultipartForm.File["file"]) == 0 {
		return
	}
	//check if the type is image
	files := r.MultipartForm.File["file"][0]
	ext := strings.ToLower(path.Ext(files.Filename))
	if ext != ".jpg" && ext != ".png" {
		return
	}
	//save img to the host
	file, err := files.Open()
	defer file.Close()
	tools.HandleError("file.open", err, -1)
	var id = "0001"                               //the temply id of goods
	savelocation := imgpath + username + id + ext //the path of it images after saving
	os.Mkdir(imgpath, os.ModePerm)
	cur, err := os.Create(savelocation)
	defer cur.Close()
	tools.HandleError("os.Create err", err, -1)
	io.Copy(cur, file)
	fmt.Println("Save file :", savelocation)
	//return the imgurl of get it iamges
	imgrul := `http://` + config.Public_addr + "/source/images?name=" + username + id + ext
	tools.WriteJson(w, imgrul)
}

//the return the images source by given imgurl
//example of imgurl:http://localhost:8090/source/images?name=testcover.jpg
func GetImages(w http.ResponseWriter, r *http.Request) {
	tools.SetHeader(w)
	if r.Method != "GET" {
		return
	}
	vars := r.URL.Query()
	name := vars["name"][0]
	fmt.Println(name)
	//find the images and return []byte
	filepath := imgpath + name
	temp, err := ioutil.ReadFile(filepath)
	if err != nil {
		tools.HandleError("GetImages readfile error :", err, 1)
		return
	}
	w.Write(temp)
	return
}
