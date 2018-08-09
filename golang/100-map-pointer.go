package main

import "fmt"

type A struct {
	ID int
}

/*
output example

1,0xc420016098,&{1}
2,0xc420016098,&{2}
3,0xc420016098,&{3}
map[1:0xc420016098 2:0xc420016098 3:0xc420016098]
333
map[2:0xc4200122c8 3:0xc4200122d0 1:0xc4200122c0]
123

*/
func main() {
	as := []A{A{1}, A{2}, A{3}}
	m := make(map[int]*A)
	for _, a := range as {
		m[a.ID] = &a
		fmt.Printf("%d,%p,%v\n", a.ID, &a, &a)
	}
	scan(m)

	for i, _ := range as {
		m[as[i].ID] = &as[i]
	}
	scan(m)
}

func scan(m map[int]*A) {
	fmt.Println(m)
	for _, a := range m {
		print(a.ID)
	}
	print("\n")
}
