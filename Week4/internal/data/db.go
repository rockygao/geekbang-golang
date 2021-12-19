package data

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var db *sql.DB

func DBInit() error {
	var err error
	db, err = sql.Open("mysql", "root:123@tcp(localhost:3306)/golang")
	if err != nil {
		log.Panicf(" 数据库连接错误，errors: %+v  ", err)
		return err
	}
	defer db.Close() //关闭数据库

	if err = db.Ping(); err != nil {
		log.Printf("数据库连接失败 , errors: %+v  ", err)
		return err
	}
	return nil
}

type User struct {
	id   int
	name string
}

type errorString struct {
	err    error
	errStr string
}

// db errors
func (e *errorString) dbErr() error {
	if e.err == sql.ErrNoRows {
		return errors.New(e.err.Error())
	} else {
		return e.err
	}
}

// 获取一条记录 By id
func GetUser(id int) (*User, error) {
	if err := DBInit(); err != nil {
		return nil, err
	}
	defer db.Close()

	var user User
	err := db.QueryRow("SELECT id,name FROM user WHERE id = ?", id).Scan(&user.id, &user.name)
	if err != nil {
		var e errorString
		e.err = err
		return nil, e.dbErr()
	}
	return &user, nil
}

// func queryRow(queryStr string) (uname string, err error) {
// 	//"用户名:密码@[连接方式](主机名:端口号)/数据库名"
// 	Database, _ := sql.Open("mysql", "root:root@(127.0.0.1:3306)/golang") // 设置连接数据库的参数
// 	err = Database.Ping()                                                 //连接数据库
// 	if err != nil {
// 		panic("数据库链接失败")
// 	}
// 	defer Database.Close()              //关闭数据库
// 	rows := Database.QueryRow(queryStr) //获取一行数据
// 	err = rows.Scan(&uname)       //将rows中的数据存到id,name中

// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			fmt.Println(errors.Wrap(err, fmt.Sprintf("Query Not Found (%s)", queryStr)))
// 		} else {
// 			fmt.Println(errors.Wrap(err, fmt.Sprintf("Query Faild (%s)", queryStr)))
// 		}
// 	}
// 	return uname, nil
// }
