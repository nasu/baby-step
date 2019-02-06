package main

import (
	"fmt"
	"reflect"
	"time"
)

type S1 struct {
	Id   string
	Name string
}

type S2 struct {
	Id        string
	Name      string
	CreatedAt time.Time
}

type CreatedAtKeeper interface {
	setCreatedAt(time.Time)
}

func (s *S2) setCreatedAt(t time.Time) {
	s.CreatedAt = t
}

func main() {
	fmt.Println(hasCreatedAtWithInterface(&S1{}))
	fmt.Println(hasCreatedAtWithInterface(&S2{}))
	fmt.Println(hasCreatedAtWithReflect(&S1{}))
	fmt.Println(hasCreatedAtWithReflect(&S2{}))
}

func hasCreatedAtWithInterface(s interface{}) bool {
	_, ok := s.(CreatedAtKeeper)
	return ok
}

func hasCreatedAtWithReflect(s interface{}) bool {
	val := reflect.ValueOf(s).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		if field.Name != "CreatedAt" {
			continue
		}
		return true
	}
	return false
}
