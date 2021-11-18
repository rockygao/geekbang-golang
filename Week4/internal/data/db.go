package data

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

// 定义一个全局对象db
var Database *sql.DB




func queryRow(queryStr string) (uname string, err error) {
	//"用户名:密码@[连接方式](主机名:端口号)/数据库名"
	Database, _ := sql.Open("mysql", "root:root@(127.0.0.1:3306)/golang") // 设置连接数据库的参数
	err = Database.Ping()                                                 //连接数据库
	if err != nil {
		panic("数据库链接失败")
	}
	defer Database.Close()              //关闭数据库
	rows := Database.QueryRow(queryStr) //获取一行数据
	err = rows.Scan(&uname)       //将rows中的数据存到id,name中

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println(errors.Wrap(err, fmt.Sprintf("Query Not Found (%s)", queryStr)))
		} else {
			fmt.Println(errors.Wrap(err, fmt.Sprintf("Query Faild (%s)", queryStr)))
		}
	}
	return uname, nil
}
