package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/h2non/filetype.v1"
	"log"
	"github.com/ghodss/yaml"
	"os"
	"./engine"
	"strconv"
)

type Config struct {
	Root string `json:"root"`
}

type LsFile struct {
	Name string `json:"name"`
	Size string `json:"size"`
	Type string `json:"type"`
}

type Return struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func GetListHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	file, err := os.Open("./config.yaml")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	data := make([]byte, 1000)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	config := Config{}
	_ = yaml.Unmarshal(data[:count], &config)


	urlObj := r.URL
	query := urlObj.Query()
	pth := query["path"]
	pathStr := config.Root
	listfile := make([]LsFile, 0)
	if len(pth) != 0 {
		pathStr = pth[0]
	}
	files, err := ioutil.ReadDir(pathStr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, v := range files {
		Type := "file"
		if v.IsDir() == true {
			Type = "directory"
		}
		unit := "B"
		size := int(v.Size())
		if size > 1024 {
			unit = "K"
			size = size >> 10
		}
		if size > 1024 {
			unit = "M"
			size = size >> 10
		}
		if size > 1024 {
			unit = "G"
			size = size >> 10
		}
		human := strconv.Itoa(size)+unit
		listfile = append(listfile, LsFile{
			Name: v.Name(),
			Size: human,
			Type: Type,
		})
	}
	rtv := Return{0, "OK", listfile}
	retval, err := json.Marshal(rtv)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(retval)
}

func HtmlHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(engine.HtmlIndex()))
}

func ScriptHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "text/javascript")
	w.Write([]byte(engine.Script(ps.ByName("filename"))))
}

func CssHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "text/css")
	w.Write([]byte(engine.Css(ps.ByName("filename"))))
}

func AssetsHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	buf, _ := ioutil.ReadFile("./assets/"+ps.ByName("filename"))
	kind, unknow := filetype.Match(buf)
	if unknow != nil {
		log.Fatal(unknow.Error())
	}
	w.Header().Set("Content-Type", kind.MIME.Value)
	w.Write([]byte(engine.Assets(ps.ByName("filename"))))
}

func main() {
	router := httprouter.New()
	router.GET("/", HtmlHandler)
	router.GET("/files", GetListHandler)
	router.GET("/dist/scripts/:filename", ScriptHandler)
	router.GET("/dist/styles/:filename", CssHandler)
	router.GET("/assets/:filename", AssetsHandler)
	log.Fatal(http.ListenAndServe(":8888", router))
}