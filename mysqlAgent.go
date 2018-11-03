package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

type stu struct {
	Username string `db:"username"`
	Password string `db:"password"`
	Department string `db:"department"`
	Email string `db:"email"`
}

func init(){
	db,err := sqlx.Open("mysql","root:199999@tcp(127.0.0.1:3306)/beego?charset=utf8")
	if err != nil {
		fmt.Println("open mysql failed:",err)
		return
	}
	DB = db
	DB = db
	DB.SetMaxOpenConns(30)
	DB.SetMaxIdleConns(15)
}
func getErr(nill interface{},msg string)  {
	if nill != nil {
		fmt.Println("get a ERROR:",msg)
		return
	}
}

func main() {
	// eg:do just crud
	//result, err := DB.Exec("insert into userinfo (username, password, department, email) values (?,?,?,?)","ZHANGLIWEN","199999","GOLANG","1014309518@QQ.COM")
	//getErr(err,"insert get errors")
	//id,_ := result.LastInsertId()
	//fmt.Println("insert id is:",id)

	//rows, errGet := DB.Query("SELECT username,password,email FROM userinfo")
	//getErr(errGet,"查询失败")
	//for rows.Next() {
	//	var username,password,email string
	//		rows.Scan(&username, &password, &email)
	//	println(username,password,email)
	//}


	//_,err1 := DB.Exec("update userinfo set username = ? where uid = ?","zhangliwen",1)
	//getErr(err1,"update get errors")
	//
	//_, err2 := DB.Exec("delete from userinfo where uid = ? ", 4)
	//getErr(err2,"删除错误")

	 students := []stu{}
	err := DB.Select(&students,"select username,password,email From userinfo")
	getErr(err,"查询失败")
	for _,student:=  range students {
		fmt.Println("student info:",student.Username,student.Password,student.Email)
	}





}
