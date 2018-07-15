package main

import (
	"io/ioutil"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"log"
	"github.com/ghodss/yaml"
	"./engine"
	"encoding/json"
)

type Config struct {
	Root string
}

func main() {
	router := httprouter.New()
	config := Config{}
	buf, err := ioutil.ReadFile("./config.yaml")
	_ = yaml.Unmarshal(buf, &config)
	if err != nil {
		log.Fatal(err.Error())
		return 
	}

	router.GET("/files", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
		query := r.URL.Query()
		pathQuery := ""
		if query["path"] != nil {
			pathQuery = query["path"][0]
		}
		paramstr := engine.FileListParams{
			Path: pathQuery,
		}
		confstr, err := json.Marshal(config)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		param, err := json.Marshal(paramstr)
		engine.FileListHandler(string(param), w, string(confstr))
	})

	router.GET("/file", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
		query := r.URL.Query()
		pathQuery := ""
		fileName := ""
		if query["path"] != nil {
			pathQuery = query["path"][0]
		}
		if query["filename"] != nil {
			fileName = query["filename"][0]
			paramstr := engine.DownloadParams{
				Path: pathQuery,
				Name: fileName,
			}
			confstr, err := json.Marshal(config)
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			param, err := json.Marshal(paramstr)
			engine.DownloadHandler(string(param), w, string(confstr))
		} else {
			paramstr := engine.ViewParams{
				Path: pathQuery,
			}
			confstr, err := json.Marshal(config)
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			param, err := json.Marshal(paramstr)
			engine.ViewHandler(string(param), w, string(confstr))
		}
	})

	log.Fatal(http.ListenAndServe(":8888", router))
}