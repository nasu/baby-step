package main

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/mohae/deepcopy"
	"github.com/ulule/deepcopier"
)

type S struct {
	ID       int
	Name     string
	ChildRef *SS
	Child    SS
}

type SS struct {
	Name string
}

func main() {
	var a, b *S
	a = newS()
	b = reference(a)
	fmt.Println("REFERENCE")
	fmt.Println("  Src", a.Name, a.ChildRef.Name, a.Child.Name) // bar, bar, bar
	fmt.Println("  Dst", b.Name, b.ChildRef.Name, b.Child.Name) // bar, bar, bar

	a = newS()
	b = builtinShallowCopy(a)
	fmt.Println("BUILTIN COPY")
	fmt.Println("  Src", a.Name, a.ChildRef.Name, a.Child.Name) // foo, bar, foo
	fmt.Println("  Dst", b.Name, b.ChildRef.Name, b.Child.Name) // bar, bar, bar

	a = newS()
	b = selfShallowCopy(a)
	fmt.Println("SELF COPY")
	fmt.Println("  Src", a.Name, a.ChildRef.Name, a.Child.Name) // foo, bar, foo
	fmt.Println("  Dst", b.Name, b.ChildRef.Name, b.Child.Name) // bar, bar, bar

	a = newS()
	b = jinzhuCopier(a)
	fmt.Println("github.com/jinzhu/copier")
	fmt.Println("  Src", a.Name, a.ChildRef.Name, a.Child.Name) // foo, bar, foo
	fmt.Println("  Dst", b.Name, b.ChildRef.Name, b.Child.Name) // bar, bar, bar

	a = newS()
	b = ululeDeepCopier(a)
	fmt.Println("github.com/ulule/deepcopier")
	fmt.Println("  Src", a.Name, a.ChildRef.Name, a.Child.Name) // foo, bar, foo
	fmt.Println("  Dst", b.Name, b.ChildRef.Name, b.Child.Name) // bar, bar, bar

	a = newS()
	b = mohaeDeepcopy(a)
	fmt.Println("github.com/mohae/deepcopy")
	fmt.Println("  Src", a.Name, a.ChildRef.Name, a.Child.Name) // foo, foo, foo
	fmt.Println("  Dst", b.Name, b.ChildRef.Name, b.Child.Name) // bar, bar, bar

	a = newS()
	b = selfDeepcopy(a)
	fmt.Println("SELF DEEPCOPY")
	fmt.Println("  Src", a.Name, a.ChildRef.Name, a.Child.Name) // foo, foo, foo
	fmt.Println("  Dst", b.Name, b.ChildRef.Name, b.Child.Name) // bar, bar, bar
}

func newS() *S {
	a := &S{
		ID:       1,
		Name:     "foo",
		ChildRef: &SS{"foo"},
		Child:    SS{"foo"},
	}
	return a
}

func changeS(a *S) {
	a.Name = "bar"
	a.ChildRef.Name = "bar"
	a.Child.Name = "bar"
}

func reference(a *S) *S {
	b := a
	changeS(b)
	return b
}

func builtinShallowCopy(a *S) *S {
	b := make([]S, 1) // 受け口を1つは用意しておかないとダメ＝0だと入ってこない
	copy(b, []S{*a})
	changeS(&b[0])
	return &b[0]
}

func selfShallowCopy(a *S) *S {
	b := *a
	changeS(&b)
	return &b
}

func selfDeepcopy(a *S) *S {
	bytes, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	b := &S{}
	if err := json.Unmarshal(bytes, b); err != nil {
		panic(err)
	}
	changeS(b)
	return b
}

func jinzhuCopier(a *S) *S {
	b := &S{}
	copier.Copy(b, a)
	changeS(b)
	return b
}

func ululeDeepCopier(a *S) *S {
	b := &S{}
	deepcopier.Copy(a).To(b)
	changeS(b)
	return b
}

func mohaeDeepcopy(a *S) *S {
	i := deepcopy.Copy(a)
	b, ok := i.(*S)
	if !ok {
		panic("Failed to cast")
	}
	changeS(b)
	return b
}
