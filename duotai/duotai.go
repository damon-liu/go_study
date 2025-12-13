package duotai

import (
	"fmt"
)

func init() {
	fmt.Println("--------- 多态Demo ---------")
	d := Dog{
		Name:  "Bingo",
		Color: "Wight",
		Type:  "Chaiqian",
	}
	d.Sleep()
	fmt.Println(d.Name + " color is: " + d.GetColor())
	fmt.Println(d.Name + " type is: " + d.GetType())
	fmt.Println()
}

type AnimalIn interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Dog struct {
	Name  string
	Color string
	Type  string
}

func (this *Dog) Sleep() {
	fmt.Println(this.Name + "is sleep")
}

func (this *Dog) GetColor() string {
	return this.Color
}

func (this *Dog) GetType() string {
	return this.Type
}
