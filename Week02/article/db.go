package article

import (
	"database/sql"
	// 引入数据库驱动注册及初始化
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/pkg/errors"
)

var (
	db *sql.DB
)

func openConn() (err error){
	db, err = sql.Open("mysql","root:root@tcp(localhost:3306)/geektime-go-week02?charset=utf8")
	if err != nil {
		fmt.Println("数据库链接错误,", err)
		return
	}
	return nil
}


func GetArticleById(id string) (*Article,error) {
	sql := "select id,title from t_article where id= ?"
	rows , err := query(sql,id)
	return rows,err
}

func query ( querySql string,id string) (*Article,error) {
	err := openConn()
	article := new(Article)
	err = db.QueryRow(querySql,id).Scan(&article.Id,&article.Title)
	if err != nil {
		if errors.Is(err,sql.ErrNoRows) {
			return nil,errors.New("query aritcle errors :" + sql.ErrNoRows.Error())
		}
		return nil,errors.Wrapf(err,"数据库访问失败")
	}
	return article,err
}

