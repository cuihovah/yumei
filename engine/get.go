package engine

import (
	"gopkg.in/h2non/filetype.v1"
	"io/ioutil"
	"path"
	"encoding/json"
	"net/http"
	"log"
	"path/filepath"
	"strings"
)

type downloadParams struct {
	Filename string
} 

type FileItem struct {
	Name string `json:"name"`
	Size int64 `json:"size"`
	Type string `json:"type"`
}

func DownloadHandler(paramstr string, w http.ResponseWriter, confstr string) {
	var params DownloadParams
	var config Config
	err := json.Unmarshal([]byte(paramstr), &params)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(confstr), &config)
	if err != nil {
		return
	}
	fullPath := config.Root
	if params.Path != "" {
		fullPath = path.Join(fullPath, params.Path)
	}
	rootPath, _ := filepath.Abs(config.Root)
	resolvePath, err := filepath.Abs(fullPath)
	returnPath := strings.Replace(resolvePath, rootPath, "", -1)

	if len(returnPath) == 0 || []byte(returnPath)[0] != '/' {
		returnPath = "/" + returnPath
	}
	if err != nil {
		w.Header().Set("Content-Type", "text/plain;charset=UTF-8")
		w.Write([]byte(err.Error()))
		return
	}
	if strings.Index(resolvePath, rootPath) != 0 {
		w.Header().Set("Content-Type", "text/plain;charset=UTF-8")
		w.WriteHeader(403)
		w.Write([]byte(returnPath+" 非法访问"))
		return
	}
	content, _ := ioutil.ReadFile(resolvePath)
	kind, unknow := filetype.Match(content)
	if unknow != nil {
		log.Fatal(unknow.Error())
	}
	w.Header().Set("Content-Type", kind.MIME.Value)
	w.Header().Set("Content-Disposition", "attachment")
	w.Header().Set("filename", params.Name)
	w.Header().Set("Content-Length", string(len(content)))
	w.Write(content)
}


func ViewHandler(paramstr string, w http.ResponseWriter, confstr string) {
	var params ViewParams
	var config Config
	err := json.Unmarshal([]byte(paramstr), &params)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(confstr), &config)
	if err != nil {
		return
	}
	fullPath := config.Root
	if params.Path != "" {
		fullPath = path.Join(fullPath, params.Path)
	}
	rootPath, _ := filepath.Abs(config.Root)
	resolvePath, err := filepath.Abs(fullPath)
	returnPath := strings.Replace(resolvePath, rootPath, "", -1)

	if len(returnPath) == 0 || []byte(returnPath)[0] != '/' {
		returnPath = "/" + returnPath
	}
	if err != nil {
		w.Header().Set("Content-Type", "text/plain;charset=UTF-8")
		w.Write([]byte(err.Error()))
		return
	}
	if strings.Index(resolvePath, rootPath) != 0 {
		w.Header().Set("Content-Type", "text/plain;charset=UTF-8")
		w.WriteHeader(403)
		w.Write([]byte(returnPath+" 非法访问"))
		return
	}
	content, _ := ioutil.ReadFile(resolvePath)
	kind, unknow := filetype.Match(content)
	if unknow != nil {
		log.Fatal(unknow.Error())
	}
	w.Header().Set("Content-Type", kind.MIME.Value)
	w.Header().Set("Content-Length", string(len(content)))
	w.Write(content)
}

func FileListHandler(paramstr string, w http.ResponseWriter, confstr string) {
	var params FileListParams
	var config Config
	err := json.Unmarshal([]byte(paramstr), &params)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(confstr), &config)
	if err != nil {
		return
	}
	fullPath := config.Root
	if params.Path != "" {
		fullPath = path.Join(fullPath, params.Path)
	}
	rootPath, _ := filepath.Abs(config.Root)
	resolvePath, err := filepath.Abs(fullPath)
	returnPath := strings.Replace(resolvePath, rootPath, "", -1)

	if len(returnPath) == 0 || []byte(returnPath)[0] != '/' {
		returnPath = "/" + returnPath
	}
	if err != nil {
		w.Header().Set("Content-Type", "text/plain;charset=UTF-8")
		w.Write([]byte(err.Error()))
		return
	}
	if strings.Index(resolvePath, rootPath) != 0 {
		w.Header().Set("Content-Type", "text/plain;charset=UTF-8")
		w.WriteHeader(403)
		w.Write([]byte(returnPath+" 非法访问"))
		return
	}
	files, err := ioutil.ReadDir(resolvePath)
	listfile := []FileItem{}
	if err != nil {
		return
	}
	for _, v := range files {
		Type := "file"
		if v.IsDir() == true {
			Type = "directory"
		}
		listfile = append(listfile, FileItem{
			Name: v.Name(),
			Size: int64(v.Size()),
			Type: Type,
		})
	}
	rtv := Return{0, "OK", struct{
		Path string `json:"path"`
		Files []FileItem `json:"files"`
	}{
		returnPath, 
		listfile,
	}}
	retval, err := json.Marshal(rtv)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(retval)
}