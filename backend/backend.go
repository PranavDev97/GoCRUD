package backend

import(
	"GoCRUD/domain"
	"GoCRUD/conf"
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"encoding/json"
)

func Create(studentDetails domain.StudentCredentials) (response string,err error) {
	var resp string
	var mysqlData conf.Vars
	var connect string
	mysqlData=conf.MysqlData
	connect=mysqlData.Username+":"+mysqlData.Password+"@/"+mysqlData.DB
	db,err := sql.Open("mysql",connect)
	if err!=nil {
		fmt.Println("db connection error:",err)
		resp="Insertion into DB failed"
		return resp,err
	}
	stmt,err := db.Prepare("INSERT INTO student(S_id,D_id,Marks) values(?,?,?)");
	if err!=nil {
		fmt.Println("statement creation failed:",err)
		resp="Insertion into DB failed"
		return resp,err
	}
	res,err := stmt.Exec(studentDetails.Sid,studentDetails.Did,studentDetails.Marks);
	fmt.Println(res)
	if err!=nil {
		fmt.Println("insertion failed:",err)
		resp="Insertion into DB failed"
		return resp,err
	}else{
		resp="Insertion into DB success"
	}
	return resp,nil
}

func Delete(sid int) (response string,err error){
	var resp string
	var mysqlData conf.Vars
	var connect string
	mysqlData=conf.MysqlData
	connect=mysqlData.Username+":"+mysqlData.Password+"@/"+mysqlData.DB
	db,err := sql.Open("mysql",connect)
	if err!=nil {
		fmt.Println("db connection error:",err)
		resp="Insertion into DB failed"
		return resp,err
	}
	stmt,err := db.Prepare("DELETE FROM student WHERE S_id=?")
	if err != nil {
        fmt.Println("Statement preperation error:",err)
		resp="Deletion from DB failed"
		return resp,err
    }
	res,err:=stmt.Exec(sid)
	fmt.Println(res)
	if err!=nil {
		fmt.Println("Delete failed:",err)
		resp="Deletion from DB failed"
		return resp,err
	}else{
		resp="Record successfuly deleted"
	}
	return resp,nil
}

func Read(sid int) (response string,err error) {
	var resp string
	var mysqlData conf.Vars
	var connect string
	mysqlData=conf.MysqlData
	connect=mysqlData.Username+":"+mysqlData.Password+"@/"+mysqlData.DB
	db,err := sql.Open("mysql",connect)
	if err!=nil {
		fmt.Println("db connection error:",err)
		resp="Read failed"
		return resp,err
	}
	stmt, err := db.Query("SELECT * FROM student WHERE S_id=?", sid)
	if err != nil {
        fmt.Println("Statement preperation error:",err)
		resp="Read failed"
		return resp,err
	}
	var studentDetails domain.StudentCredentials
	stmt.Next();
	er:=stmt.Scan(&studentDetails.Sid,&studentDetails.Did,&studentDetails.Marks)
	if er != nil {
        fmt.Println("Statement execution error:",er)
		resp="Read failed"
		return resp,err
	}
	studentData, err := json.Marshal(studentDetails)
	return string(studentData),nil
}

func Update(sid int,marks int) (response string,err error){
	var resp string
	var mysqlData conf.Vars
	var connect string
	mysqlData=conf.MysqlData
	connect=mysqlData.Username+":"+mysqlData.Password+"@/"+mysqlData.DB
	db,err := sql.Open("mysql",connect)
	if err!=nil {
		fmt.Println("db connection error:",err)
		resp="Update failed"
		return resp,err
	}
	stmt, err := db.Prepare("UPDATE student SET Marks=? WHERE S_id=?")
	if err != nil {
        fmt.Println("Statement preperation error:",err)
		resp="Update failed"
		return resp,err
	}
	res,err:=stmt.Exec(marks,sid)
	fmt.Println(res)
	if err != nil {
        fmt.Println("Statement Execution error:",err)
		resp="Update failed"
		return resp,err
	}
	resp="Updation Successful"
	return resp,nil
}