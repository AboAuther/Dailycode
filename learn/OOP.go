package main

import "fmt"

type BaseBird struct {
	age int
}

func (this *BaseBird) Cal() {
	this.Add()
}
func (this *BaseBird) Add() {
	fmt.Printf("before add: age=%d\n", this.age)
	this.age = this.age + 1
	fmt.Printf("after add: age=%d\n", this.age)
}

type DerivedBird struct {
	BaseBird
}

func (this *DerivedBird) Add() {
	fmt.Printf("before add: age=%d\n", this.age)
	this.age = this.age + 2
	fmt.Printf("after add: age=%d\n", this.age)
}

func main() {
	var b1 BaseBird
	var b2 DerivedBird
	b1 = BaseBird{age: 1}
	b1.Cal()
	b2 = DerivedBird{BaseBird{1}}
	b2.Cal()
}
