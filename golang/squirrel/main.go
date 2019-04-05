package squirrel

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

func InsertSample() {
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
	insert = insert.Columns("CreatedAt")
	values = append(values, sq.Expr("CURRENT_TIMESTAMP()"))
	insert = insert.Values(values...)
	sql, params, err := insert.ToSql()
	if err != nil {
		panic(err)
	}
	fmt.Println(sql, params)
}

func UseSpannerFormat() {
	f := new(spannerFormat)
	f.named = []string{"col_a", "col_b"}
	sql, params, err := sq.StatementBuilder.
		PlaceholderFormat(f).
		Select("*").From("table").
		Where(sq.Eq{"col_a": 1}).
		Where(sq.Eq{"col_b": 2}).
		ToSql()
	if err != nil {
		panic(err)
	}
	fmt.Println(sql, params)
}
