package main

import (
	"fmt"
)

type A struct {
	name string
}

type B struct {
	name string
}

type C struct {
	name string
	a    A
}

type IA interface {
	a()
}

type IB interface {
	a()
}

func (a *A) a() {

}
func (b *B) a() {

}

func main() {
	cp(1, A{"你"}, A{"你"})
	cp(2, &A{"你"}, &A{"你"})

	cp(3, A{"A"}, B{"B"})

	cp(4, C{"C", A{"A"}}, C{"C", A{"A"}})
	cp(5, &C{"C", A{"A"}}, &C{"C", A{"A"}})

	var ia1 IA
	var ia2 IA
	var ib1 IB
	cp(6, ia1, ia2)
	ia1 = func() *A { return nil }()
	cp(7, ia1, ia2)
	cp("7-1", ia1, nil)
	ia2 = func() IA { return nil }()
	cp(8, ia1, ia2)
	ia2 = (IA)(nil)
	cp(9, ia1, ia2)

	ia1 = IA(nil)
	ib1 = IB(nil)
	cp(10, ia1, ib1)

	s := &A{"A"}
	ia1 = s
	ib1 = s
	cp(11, ia1, ib1)

	ib1 = IB(s)
	cp(12, ia1, ib1)
}

func cp(i interface{}, a interface{}, b interface{}) {
	fmt.Printf("--%v--\n  a:%T%+v\n  b:%T%+v\n  结果：%v\n\n", i, a, a, b, b, a == b)
}
