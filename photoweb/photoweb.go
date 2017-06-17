package main

import (
	"io"
	"log"
	"html/template"
	"net/http"
	"os"
	"io/ioutil"
	"fmt"
)
 const(
	 UPLOAD_DIR ="./uploads"
 )
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.Method == "GET" {
		   t, _ := template.ParseFiles("upload.html") //加载html
			log.Println(t.Execute(w, nil)) 
		return
	}
	if r.Method=="POST"{
		f,h,err:=r.FormFile("image")//此值与html的取值有关
		fmt.Println(err)
		if err!=nil {
			http.Error(w, err.Error(),http.StatusInternalServerError)
			return
		}
	filename:=h.Filename
	defer f.Close()
	t,err:=os.Create(UPLOAD_DIR+"/"+filename)
/*
*/
	fmt.Println(t)
	if err!=nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer t.Close()
		if _,err:=io.Copy(t,f); err!=nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/View?id="+filename,http.StatusInternalServerError)
	}
}
func viewHandle(w http.ResponseWriter,r *http.Request){

	imageId :=r.FormValue("id")
	imagePath:=UPLOAD_DIR +"/" +imageId
	if exists :=isExists(imagePath) ; !exists{
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r,imagePath)

}
func isExists(path string) bool{
	_,err:=os.Stat(path)
	if err!=nil{
		return os.IsExist(err)
	}
	return true
}
func listHandler(w http.ResponseWriter,r *http.Request)  {
	fileInfoArr,err:=ioutil.ReadDir("./uploads")
	if err !=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return 
	}
	locals :=make(map[string] interface{})
	images :=[]string{}
	for _,fileInfo:=range fileInfoArr{
		images =append(images,fileInfo.Name())
	}
	locals["images"]=images
	t,err:=template.ParseFiles("list.html")
	if err !=nil{
		http.Error(w, err.Error(),http.StatusInternalServerError)
		return
	}
	t.Execute(w, locals)
}
func main() {

	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/view", viewHandle)
	http.HandleFunc("/",listHandler)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err.Error())
	}
}

