package main

import "fmt"

type IAnimal interface {
	bark() string
}

type Dog struct{}

func (Dog) bark() string {
	return "woof"
}

type BullDog struct {
	d    Dog
	name string
}

func (bd *BullDog) bark() string {
	return fmt.Sprintf("%s: %s", bd.name, bd.d.bark())
}

type Cat struct{}

func (Cat) bark() string {
	return "meow"
}

func PrintBark(ani IAnimal) {
	fmt.Println(ani.bark())
}

func main() {
	d := Dog{}
	c := Cat{}
	bd := BullDog{d: Dog{}, name: "sam"}
	PrintBark(&d)
	PrintBark(&c)
	PrintBark(&bd)
}
