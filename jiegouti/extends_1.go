package jiegouti

import "fmt"

type Cat struct {
	Name string
	Sex  string
}

type Tom struct {
	Cat Cat
	Say bool
}

func init() {
	// fmt.Println("---------继承Demo ---------")
	// a := Cat{
	// 	Name: "Cat",
	// 	Sex:  "xxx",
	// }
	// a.Walk()
	// a.Eat()

	// tom := Tom{
	// 	Cat: Cat{
	// 		Name: "Tom",
	// 		Sex:  "Male",
	// 	},
	// 	Say: true,
	// }
	// tom.Walk()
	// tom.Eat()
	// fmt.Println()
}

func (this *Cat) Walk() {
	fmt.Println("Anmail " + this.Name + " Walk()...")
}

func (this *Cat) Eat() {
	fmt.Println("Animal " + this.Name + " Eat()...")
}

func (this *Tom) Eat() {
	fmt.Println("Cat " + this.Cat.Name + " Eat()...")
}

func (this *Tom) Walk() {
	fmt.Println("Cat " + this.Cat.Name + " Walk()...")
}

type MiaoMiao struct {
	Cat Cat
	Say string
}

func (this *MiaoMiao) Eat() {
	fmt.Println("Cat " + this.Cat.Name + " Eat()...")
}

func (this *MiaoMiao) Walk() {
	fmt.Println("Cat " + this.Cat.Name + " Walk()...")
}
