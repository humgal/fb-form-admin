package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	mysql "fubon.com/form/server/sql"
	"fubon.com/form/server/util"
	"github.com/golang-jwt/jwt"
)

type resp struct {
	Code int
	Msg  string
}

var jwtSecret = []byte("my_jwt_secret")

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
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
	http.HandleFunc("/login", login)
	http.HandleFunc("/form/create", CreateForm)
	http.HandleFunc("/form/list", ListForm)
	http.HandleFunc("/form/detail", PostForm)
	http.HandleFunc("/form/updatecontent", UpdateFormContent)
	http.HandleFunc("/form/hislist", ListFormHis)
	http.HandleFunc("/form/hisdetail", HisDetail)
	http.ListenAndServe(util.Config.Address, nil)
}

func login(writer http.ResponseWriter, request *http.Request) {
	var user User
	var res resp
	b, err := ioutil.ReadAll(request.Body)
	if err != nil {
		util.Logger.Println(err.Error())

		return
	}
	util.Logger.Println(string(b))

	err = json.Unmarshal(b, &user)
	if err != nil {
		util.Logger.Println(err.Error())

		return
	}

	DB, err := sql.Open(mysql.DbDriverName, mysql.DbName)
	if err != nil {
		util.Logger.Println(err.Error())

		return
	}
	match, err := mysql.UserMatch(DB, mysql.User{0, user.Name, "", user.Password, 0})
	if err != nil {
		util.Logger.Println(err.Error())
		return
	}
	if !match {
		res = resp{0, "用户或密码错误"}
		writer.Header().Set("Content-Type", "application/json")
		output, err := json.MarshalIndent(&res, "", "")
		if err != nil {
			return
		}
		writer.Write(output)
		util.Logger.Println(res)
		return
	}
	token, err := GenerateToken(user.Name, user.Password)
	if err != nil {
		util.Logger.Println(err.Error())
		return
	}
	res = resp{1, token}
	writer.Header().Set("Content-Type", "application/json")
	output, err := json.MarshalIndent(&res, "", "")
	if err != nil {
		return
	}
	writer.Write(output)
	util.Logger.Println(res)
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
	token := r.Header.Get("token")
	util.Logger.Println(token)
	var cla *Claims
	cla, err := ParseToken(token)
	util.Logger.Println(cla.Username)
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
	token := r.Header.Get("token")
	util.Logger.Println(token)
	var cla *Claims
	cla, err := ParseToken(token)
	if err != nil {
		util.Logger.Println(err)
	}
	util.Logger.Println(cla.Username)

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
	mysqlformcontent.UpName = cla.Username
	mysqlformcontent.UpTime = time.Now().Format("2006-01-02 15:04:05")
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
		res.Code = 1
	}
	var conetent map[string]interface{}
	err = json.Unmarshal([]byte(formcontent.Content), &conetent)
	if err != nil {
		util.Logger.Println(err)
		res.Code = 1
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
					switch value.(type) {

					case string:
						in, _ := strconv.Atoi(value.(string))
						formRules[i].Children = append(formRules[i].Children, formRules[i].Options[in-1].Lable)

					case []interface{}:
						valslice := value.([]interface{})
						for j := 0; j < len(valslice); j++ {
							in, _ := strconv.Atoi(valslice[j].(string))
							formRules[i].Children = append(formRules[i].Children, formRules[i].Options[in-1].Lable)
							formRules[i].Children = append(formRules[i].Children, ",")
						}
					default:
						valslice := value.([]string)
						println(valslice)
					}

				} else {
					formRules[i].Children = append(formRules[i].Children, value.(string))
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
	var lastrule FormRule
	lastrule.Children = append(lastrule.Children, formcontent.UpTime)
	lastrule.Type = "span"
	lastrule.Tag = "span"
	lastrule.Title = "上传时间"
	lastrule.Display = true
	lastrule.Hidden = false
	lastrule.Native = false
	lastrule.Info = ""
	lastrule.Value = ""
	resrules = append(resrules, lastrule)

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

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * 24 * time.Hour)

	claims := Claims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{

			ExpiresAt: expireTime.Unix(),

			Issuer: "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err

}
