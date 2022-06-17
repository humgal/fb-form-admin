package mysql

import (
	"database/sql"
	"errors"

	"fubon.com/form/server/util"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DbDriverName = "mysql"
	DbName       = "root:password@(127.0.0.1:3306)/form"
)

type Form struct {
	Id     int
	Title  string
	Rule   string
	Option string
	Uper   int
	UpName string
	UpTime string
	DepId  int
	Status int
}

type FormContent struct {
	Id      int
	Title   string
	Rule    string
	Option  string
	Content string
	Uper    int
	UpName  string
	UpTime  string
	DepId   int
	Status  int
	FormId  int
}

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
	DepId    int
}

type Department struct {
	Id   int
	Name string
	Role int
}

type UserUnionDep struct {
	UserId     int
	Name       string
	Email      string
	DepId      int
	Department string
	Role       int
}

func UserMatch(db *sql.DB, user User) (bool, error) {
	sql := `select count(id) from users where name=? and password=?`

	stmt, err := db.Prepare(sql)
	if err != nil {
		util.Logger.Println(err.Error())
		return false, err
	}
	rows, err := stmt.Query(user.Name, user.Password)
	if err != nil {
		return false, err
	}
	var res = make([]int, 0)
	for rows.Next() {
		var id int
		rows.Scan(&id)
		res = append(res, id)
	}
	if res[0] >= 1 {
		return true, err
	} else {

		return false, err
	}

}

func FindUserByName(db *sql.DB, username string) (un UserUnionDep, err error) {
	sql := `select a.id, a.name,a.email,b.id,b.name,b.role from  deparment as b left join users as a on a.depid = b.id  where a.name=? `

	stmt, err := db.Prepare(sql)
	if err != nil {
		util.Logger.Println(err.Error())
		return un, err
	}
	rows, err := stmt.Query(username)
	if err != nil {
		return un, err
	}
	var res = make([]UserUnionDep, 0)
	for rows.Next() {
		var uns UserUnionDep
		rows.Scan(&uns.UserId, &uns.Name, &uns.Email, &uns.DepId, &uns.Department, &uns.Role)
		res = append(res, uns)
	}
	if len(res) == 1 {
		return res[0], err
	} else {
		return un, errors.New("没有找到该用户。。。")
	}

}

func FindAllDepartment(db *sql.DB, sysid int) (deparments []Department, err error) {
	sql := `select * from  deparment  where sysid=? `

	stmt, err := db.Prepare(sql)
	if err != nil {
		util.Logger.Println(err.Error())
		return nil, err
	}
	rows, err := stmt.Query(sysid)
	if err != nil {
		return nil, err
	}

	var res = make([]Department, 0)
	for rows.Next() {
		var dep Department
		rows.Scan(&dep.Id, &dep.Name, nil)
		res = append(res, dep)
	}
	return res, nil
}

func FindDepidByName(db *sql.DB, depname string) (dep Department, err error) {
	sql := `select * from  deparment  where name=? `

	stmt, err := db.Prepare(sql)
	if err != nil {
		util.Logger.Println(err.Error())
		return dep, err
	}
	rows, err := stmt.Query(depname)
	if err != nil {
		return dep, err
	}

	var res = make([]Department, 0)
	for rows.Next() {
		var dep Department
		rows.Scan(&dep.Id, &dep.Name, &dep.Role)
		res = append(res, dep)
	}
	if len(res) > 0 {
		return res[0], err
	} else {
		return dep, errors.New("没有找到该部门。。。")
	}

}

func FindAllForm(db *sql.DB) (forms []Form, err error) {
	sql := `select * from  form  `

	stmt, err := db.Prepare(sql)
	if err != nil {
		util.Logger.Println(err.Error())
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var res = make([]Form, 0)
	for rows.Next() {
		var form Form
		rows.Scan(&form.Id, &form.Title, &form.Rule, &form.Option, &form.Uper, &form.UpName, &form.UpTime, &form.DepId, &form.Status)
		res = append(res, form)
	}
	return res, nil
}

func InsertForm(db *sql.DB, form Form) (err error) {
	sql := " INSERT INTO `form` ( `title`, `rule`, `option`, `uper`, `upname`, `uptime`, `depid`, `status`)  values(?,?,?,?,?,?,?,?)"

	stmt, err := db.Prepare(sql)
	if err != nil {
		util.Logger.Println(err.Error())
		return err
	}
	_, err = stmt.Exec(form.Title, form.Rule, form.Option, form.Uper, form.UpName, form.UpTime, form.DepId, form.Status)
	return err
}

func InsertOrUpdateForm(db *sql.DB, form Form) (err error) {
	sql := `insert or replace into form (id,title,rule,option,uper,upname,uptime,depid,status) values (?,?,?,?,?,?,?,?,?)`

	stmt, err := db.Prepare(sql)
	if err != nil {
		util.Logger.Println(err.Error())
		return err
	}
	_, err = stmt.Exec(form.Id, form.Title, form.Rule, form.Option, form.Uper, form.UpName, form.UpTime, form.DepId, form.Status)
	return err
}

func FindFormById(db *sql.DB, id int) (form Form, err error) {
	sql := `select * from  form  where id=? `

	stmt, err := db.Prepare(sql)
	if err != nil {
		util.Logger.Println(err.Error())
		return form, err
	}
	rows, err := stmt.Query(id)
	if err != nil {
		return form, err
	}

	var res = make([]Form, 0)
	for rows.Next() {
		var form Form
		rows.Scan(&form.Id, &form.Title, &form.Rule, &form.Option, &form.Uper, &form.UpName, &form.UpTime, &form.DepId, &form.Status)
		res = append(res, form)
	}
	if len(res) > 0 {
		return res[0], nil
	} else {
		return form, errors.New("没有此条数据")
	}

}

func FindAllFormContent(db *sql.DB) (forms []FormContent, err error) {
	sql := `select * from  formcontent  `

	stmt, err := db.Prepare(sql)
	if err != nil {
		util.Logger.Println(err.Error())
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var res = make([]FormContent, 0)
	for rows.Next() {
		var form FormContent
		rows.Scan(&form.Id, &form.Title, &form.Rule, &form.Option, &form.Uper, &form.UpName, &form.UpTime, &form.DepId, &form.Status, &form.Content, &form.FormId)
		res = append(res, form)
	}
	return res, nil
}

func FindFormContentById(db *sql.DB, id int) (form FormContent, err error) {
	sql := `select * from  formcontent  where id=? `

	stmt, err := db.Prepare(sql)
	if err != nil {
		util.Logger.Println(err.Error())
		return form, err
	}
	rows, err := stmt.Query(id)
	if err != nil {
		return form, err
	}

	var res = make([]FormContent, 0)
	for rows.Next() {
		var form FormContent
		rows.Scan(&form.Id, &form.Title, &form.Rule, &form.Option, &form.Uper, &form.UpName, &form.UpTime, &form.DepId, &form.Status, &form.Content, &form.FormId)
		res = append(res, form)
	}
	if len(res) > 0 {
		return res[0], nil
	} else {
		return form, errors.New("没有此条数据")
	}
}

func InsertFormContent(db *sql.DB, form FormContent) (err error) {
	sql := "INSERT INTO `formcontent` (`title`, `rule`, `option`, `uper`, `upname`, `uptime`, `depid`, `status`, `content`, `formid`) values(?,?,?,?,?,?,?,?,?,?)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		util.Logger.Println(err.Error())
		return err
	}
	_, err = stmt.Exec(form.Title, form.Rule, form.Option, form.Uper, form.UpName, form.UpTime, form.DepId, form.Status, form.Content, form.FormId)
	if err != nil {
		util.Logger.Println(err.Error())
		return err
	}
	return nil
}

func GetRecvMailUser(db *sql.DB) (users []User, err error) {
	selsql := `select a.id,a.name,a.email,a.depid from users as a left join deparment as b on a.depid=b.id where b.role = 3  `
	stmt, err := db.Prepare(selsql)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	var res = make([]User, 0)
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Name, &user.Email, &user.DepId)
		res = append(res, user)
	}
	return res, err
}

func GetAlterMailUser(db *sql.DB) (users []User, err error) {
	selsql := `select a.id,a.name,a.email,a.depid from users as a left join deparment as b on a.depid=b.id where b.role < 3  `
	stmt, err := db.Prepare(selsql)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	var res = make([]User, 0)
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Name, &user.Email, &user.DepId)
		res = append(res, user)
	}
	return res, err
}
