package restgo

import (
	"strconv"
	"time"
	_ "github.com/go-sql-driver/mysql"

	"html/template"
	
)

var restFuncMap template.FuncMap = make(template.FuncMap)

func init(){
	restFuncMap["ctxpath"]=ctxpath
	restFuncMap["pageurl"]=pageurl
	restFuncMap["apiurl"]=apiurl
	restFuncMap["version"]=version
	restFuncMap["hello"]=hello
	restFuncMap["asset"]=asset
}

func asset() string{
	cfg := GetCfg()

	return  cfg.App["asset"]


}

func GetFuncMap()(template.FuncMap){
	return restFuncMap
}


func hello(d string) string{
	return "hello "+d
}


func ctxpath() string{
	cfg := GetCfg()
	url := cfg.App["protocal"]+"://" + cfg.App["domain"]
	if (cfg.App["port"]!="80"){
		url += (":"+cfg.App["port"])
	}
	return url
}

func pageurl(uri string) string{
	cfg := GetCfg()
	url := cfg.App["protocal"]+"://" + cfg.App["domain"]
	if (cfg.App["port"]!="80"){
		url += (":"+cfg.App["port"])
	}
	return url+"/"+uri +".shtml"
}

func apiurl(uri string) string{
	return uri
}


func version() string{
	cfg := GetCfg()
	if len(cfg.App["version"])==0{
		return  strconv.FormatInt(time.Now().Unix(),10)
	}else{
		return  cfg.App["version"]
	}

}
