package main

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

func main() {
	singerId := 1
	firstName := "Taro"
	lastName := "Hakase"
	insert := sq.Insert("Singers")
	insert = insert.Columns("SingerId")
	values := make([]interface{}, 0)
	values = append(values, sq.Expr("@SingerId", singerId))
	if firstName != "" {
		insert = insert.Columns("FirstName")
		values = append(values, sq.Expr("@FirstName", firstName))
	}
	if lastName != "" {
		insert = insert.Columns("LastName")
		values = append(values, sq.Expr("@LastName", lastName))
	}
	insert = insert.Values(values...)
	sql, params, err := insert.ToSql()
	if err != nil {
		panic(err)
	}
	fmt.Println(sql, params)
}
