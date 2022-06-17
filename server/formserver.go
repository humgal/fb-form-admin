package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	mysql "fubon.com/form/server/sql"
	"fubon.com/form/server/util"
)

type resp struct {
	Code int
	Msg  string
}

type FormRule struct {
	Type     string   `json:"type"`
	Field    string   `json:"field"`
	Title    string   `json:"title"`
	Info     string   `json:"info"`
	Native   bool     `json:"native"`
	Hidden   bool     `json:"hidden"`
	Display  bool     `json:"display"`
	Value    string   `json:"value"`
	Props    Props    `json:"props"`
	Tag      string   `json:"_fc_drag_tag"`
	Children []string `json:"children"`
	Options  []Option `json:"options"`
}

type Option struct {
	Value string `json:"value"`
	Lable string `json:"label"`
}

type Props struct {
	Title       string `json:"title"`
	Type        string `json:"type"`
	Effect      string `json:"effect"`
	Description string `json:"description"`
	Center      bool   `json:"center"`
}

type FormOne struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Upuser string `json:"upuser"`
	Uptime string `json:"uptime"`
}

type FormContent struct {
	Id      string                 `json:"id"`
	Content map[string]interface{} `json:"content"`
	Title   string                 `json:"title"`
}

type FormId struct {
	Id string `json:"id"`
}

func main() {
	util.LoadConfig()
	files := http.FileServer(http.Dir(util.Config.StaticDir))

	http.Handle("/", files)
	http.HandleFunc("/form/create", CreateForm)
	http.HandleFunc("/form/list", ListForm)
	http.HandleFunc("/form/detail", PostForm)
	http.HandleFunc("/form/updatecontent", UpdateFormContent)
	http.HandleFunc("/form/hislist", ListFormHis)
	http.HandleFunc("/form/hisdetail", HisDetail)
	http.ListenAndServe(util.Config.Address, nil)
}

func CreateForm(w http.ResponseWriter, r *http.Request) {
	formjson, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.Logger.Println(err)
	} else {
		util.Logger.Println(string(formjson))
	}
	var formRule []FormRule
	err = json.Unmarshal(formjson, &formRule)
	//store formjson
	DB, err := sql.Open(mysql.DbDriverName, mysql.DbName)
	res := resp{0, ""}
	if err != nil {
		util.Logger.Println(err)
		res.Code = 1
	}
	if len(formRule) > 0 {
		form := mysql.Form{0, formRule[0].Children[0], string(formjson), "", 1, "upname", "uptime", 1, 0}
		err = mysql.InsertForm(DB, form)
	} else {
		res.Code = 1
	}

	if err != nil {
		util.Logger.Println(err)
		res.Code = 1
	}

	w.Header().Set("Content-Type", "application/json")
	output, err := json.MarshalIndent(&res, "", "")
	if err != nil {
		return
	}
	w.Write(output)

}

func ListForm(w http.ResponseWriter, r *http.Request) {
	DB, err := sql.Open(mysql.DbDriverName, mysql.DbName)
	res := resp{0, ""}
	if err != nil {
		util.Logger.Println(err)
		res.Code = 1
	}
	forms, err := mysql.FindAllForm(DB)
	if err != nil {
		util.Logger.Println(err)
		res.Code = 1
	}
	var formList []FormOne
	for i := 0; i < len(forms); i++ {
		var formone FormOne
		formone.Id = forms[i].Id
		formone.Title = forms[i].Title
		formone.Upuser = forms[i].UpName
		formone.Uptime = forms[i].UpTime
		formList = append(formList, formone)
	}
	msg, err := json.Marshal(&formList)
	if err != nil {
		util.Logger.Println(err)
		res.Code = 1
	}
	res.Msg = string(msg)

	w.Header().Set("Content-Type", "application/json")
	output, err := json.MarshalIndent(&res, "", "")
	if err != nil {
		return
	}
	w.Write(output)
}

func PostForm(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body) //获取post的数据
	var formid FormId
	err = json.Unmarshal(b, &formid)
	res := resp{0, ""}
	if err != nil {
		util.Logger.Println(err)
	}
	DB, err := sql.Open(mysql.DbDriverName, mysql.DbName)
	if err != nil {
		util.Logger.Println(err)
	}
	id, err := strconv.Atoi(formid.Id)
	form, err := mysql.FindFormById(DB, id)
	if err != nil {
		util.Logger.Println(err)
	}
	res.Msg = form.Rule

	w.Header().Set("Content-Type", "application/json")
	output, err := json.MarshalIndent(&res, "", "")
	if err != nil {
		return
	}
	w.Write(output)
}

func UpdateFormContent(w http.ResponseWriter, r *http.Request) {
	formjson, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.Logger.Println(err)
	} else {
		util.Logger.Println(string(formjson))
	}
	var formContent FormContent
	err = json.Unmarshal(formjson, &formContent)
	//store formjson
	DB, err := sql.Open(mysql.DbDriverName, mysql.DbName)
	res := resp{0, ""}
	if err != nil {
		util.Logger.Println(err)
		res.Code = 1
	}
	var mysqlformcontent mysql.FormContent
	formid, err := strconv.Atoi(formContent.Id)
	if err != nil {
		util.Logger.Println(err)
		res.Code = 1
	}
	mysqlformcontent.FormId = formid
	mysqlformcontent.Title = formContent.Title
	mysqlformcontent.UpName = "upname"
	mysqlformcontent.UpTime = "uptime"
	contentstr, err := json.Marshal(&formContent.Content)
	mysqlformcontent.Content = string(contentstr)
	err = mysql.InsertFormContent(DB, mysqlformcontent)

	if err != nil {
		util.Logger.Println(err)
		res.Code = 1
	}

	w.Header().Set("Content-Type", "application/json")
	output, err := json.MarshalIndent(&res, "", "")
	if err != nil {
		return
	}
	w.Write(output)
}

func ListFormHis(w http.ResponseWriter, r *http.Request) {
	DB, err := sql.Open(mysql.DbDriverName, mysql.DbName)
	res := resp{0, ""}
	if err != nil {
		util.Logger.Println(err)
		res.Code = 1
	}
	forms, err := mysql.FindAllFormContent(DB)
	if err != nil {
		util.Logger.Println(err)
		res.Code = 1
	}
	var formList []FormOne
	for i := 0; i < len(forms); i++ {
		var formone FormOne
		formone.Id = forms[i].Id
		formone.Title = forms[i].Title
		formone.Upuser = forms[i].UpName
		formone.Uptime = forms[i].UpTime
		formList = append(formList, formone)
	}
	msg, err := json.Marshal(&formList)
	if err != nil {
		util.Logger.Println(err)
		res.Code = 1
	}
	res.Msg = string(msg)

	w.Header().Set("Content-Type", "application/json")
	output, err := json.MarshalIndent(&res, "", "")
	if err != nil {
		return
	}
	w.Write(output)

}

func HisDetail(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body) //获取post的数据
	var formid FormId
	err = json.Unmarshal(b, &formid)
	res := resp{0, ""}
	if err != nil {
		util.Logger.Println(err)
	}
	DB, err := sql.Open(mysql.DbDriverName, mysql.DbName)
	if err != nil {
		util.Logger.Println(err)
	}
	id, err := strconv.Atoi(formid.Id)
	formcontent, err := mysql.FindFormContentById(DB, id)
	if err != nil {
		util.Logger.Println(err)
	}
	var conetent map[string]string
	err = json.Unmarshal([]byte(formcontent.Content), &conetent)
	if err != nil {
		util.Logger.Println(err)
	}
	form, err := mysql.FindFormById(DB, formcontent.FormId)
	if err != nil {
		util.Logger.Println(err)
		res.Code = 1
	}
	var formRules []FormRule
	err = json.Unmarshal([]byte(form.Rule), &formRules)

	for key, value := range conetent {
		for i := 0; i < len(formRules); i++ {
			if formRules[i].Field == key {
				if len(formRules[i].Options) > 0 {
					for j := 0; j < len(formRules[i].Options); j++ {
						if formRules[i].Options[j].Value != value {
							formRules[i].Children = append(formRules[i].Children, formRules[i].Options[j].Lable)
						}
					}
				} else {
					formRules[i].Children = append(formRules[i].Children, value)
				}

				formRules[i].Type = "span"
				formRules[i].Tag = "span"
				formRules[i].Display = true
				formRules[i].Hidden = false
				formRules[i].Native = false
				formRules[i].Info = ""
				formRules[i].Value = ""

			}
		}
	}
	var resrules []FormRule
	for i := 0; i < len(formRules); i++ {
		if formRules[i].Type == "span" {
			resrules = append(resrules, formRules[i])
		}
	}

	Msg, err := json.Marshal(resrules)
	if err != nil {
		util.Logger.Println(err)
		res.Code = 1
	}
	res.Msg = string(Msg)

	w.Header().Set("Content-Type", "application/json")
	output, err := json.MarshalIndent(&res, "", "")
	if err != nil {
		return
	}
	w.Write(output)
}
