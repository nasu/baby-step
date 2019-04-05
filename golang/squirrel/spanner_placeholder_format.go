package squirrel

import (
	"bytes"
	"fmt"
	"strings"
)

type spannerFormat struct {
	named []string
}

const prefix = "@"

func (f spannerFormat) ReplacePlaceholders(sql string) (string, error) {
	buf := new(bytes.Buffer)
	var i int
	for {
		p := strings.Index(sql, "?")
		if p == -1 {
			break
		}
		// ?? の場合の置き換えしてあげる必要あり
		buf.WriteString(sql[:p])
		//fmt.Println(buf.String())
		fmt.Fprintf(buf, "%s%s", prefix, f.named[i])
		i++
		sql = sql[p+1:]
	}
	buf.WriteString(sql)
	return buf.String(), nil
}
