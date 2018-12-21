package conrtrol

import (
	"strconv"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fenye/model"
	"fenye/util"
)

func Viewdemo(w http.ResponseWriter, r *http.Request) {
	buf, _ := ioutil.ReadFile("view/demo.html")
	w.Write(buf)
}

//显示分页的逻辑操作
func ArtiPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var pi = r.FormValue(`pi`)
	pii,_: = strconv.Atoi(pi)
	
	var ps = r.FormValue(`ps`)
	pss,_ := strconv.Atoi(ps)
	mod,err := model.ArticlPage(pii,pss)
	if err != nil {
		w.Write(util.NewResult(300,"信息错误"))
		return
	}
	md, _ := json.Marshal(mod)
	w.Header().Set(`Content-Type`, `application/json`)
	w.Write(util.NewPage(200,"正常",md,count))
	
}