package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type QueryError struct {
	Err   error
	query string
}

//数据查询函数
func query(query string) error {
	err := &QueryError{
		Err:   sql.ErrNoRows,
		query: query,
	}
	//将执行失败的原因、携带执行语句以及执行事件一并返回给函数调用方,便于调用方进行后续其他业务逻辑处理
	err_s := fmt.Errorf("sql执行语句为: %s\n执行报错原始信息是：%w\n执行时间：%v\n", err.query, err.Err,time.Now().Unix())
	return err_s
}

func main() {
	sql_exec := "select name from user_info"
	err := query(sql_exec)
	fmt.Println(err)
	if errors.Is(errors.Unwrap(err), sql.ErrNoRows) {
		fmt.Println(errors.Unwrap(err))
	} else {
		fmt.Println("This error is not sql errors,the details: %v", err)
	}
}
