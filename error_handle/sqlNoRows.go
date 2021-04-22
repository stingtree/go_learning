package main

import (
	sql2 "database/sql"
	"fmt"
	"github.com/pkg/errors"
)

type QueryRows struct {
	query string
	err   error
}

func query(query string) error {
	sql := &QueryRows{
		query: query,
		err:   sql2.ErrNoRows,
	}
	err := errors.WithMessagef(sql.err, "NoRowsFound with %s", sql.query)
	return err
}

func main() {
	query_string := "select name from user_info"
	err := query(query_string)
	//print stack,lock the error occurred line
	fmt.Printf("%+v \n", errors.WithStack(err))
	err_cause := errors.Cause(err)
	//measure affairs according to error cause
	switch err_cause {
	case sql2.ErrNoRows:
		println("NoRows")
	default:
		println("Learing Go")
	}
}
