package main

import "fmt"

/*
comparator
comparable
enumerable
*/

type A interface {
	MMethod()
}

type AA interface {
	A
}

type B interface {
	SMethod()
	~int | uint
}

type BB interface {
	B
}

func C(p A)   {}
func CC(p AA) {}
func CCC[T B](p T) T {
	return p
}
func CCCC[T BB](p T) T {
	return p
}

type MyInt int

func (m MyInt) SMethod() {
	fmt.Println("hello")
}

func main() {
	var a MyInt = 10
	r := CCCC(a)
	fmt.Println(r)
}
