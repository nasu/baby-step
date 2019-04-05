package main

import (
	"fmt"

	"github.com/Masterminds/squirrel"
)

func main() {
	a := squirrel.Select("*").Column(squirrel.Alias(squirrel.Expr("hoge"), "b")).From("Table")
	fmt.Println(a.ToSql())
}
