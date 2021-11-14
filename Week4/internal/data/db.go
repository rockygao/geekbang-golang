package data

//我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，
//是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

// func main() {
// 	//获取用户名
// 	sqlStr := "select name from user where id=1"
// 	res, err := queryRow(sqlStr)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(res)
// 	}

// }

func queryRow(queryStr string) (uname string, err error) {
	//"用户名:密码@[连接方式](主机名:端口号)/数据库名"
	db, _ := sql.Open("mysql", "root:root@(127.0.0.1:3306)/golang") // 设置连接数据库的参数
	err = db.Ping()                                                 //连接数据库
	if err != nil {
		panic("数据库链接失败")
	}
	defer db.Close()              //关闭数据库
	rows := db.QueryRow(queryStr) //获取一行数据
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
