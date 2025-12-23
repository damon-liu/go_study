package jiegouti

import "fmt"

func init() {
	// extendsDemno()
}

func ExtendsDemno() {
	d := dog{
		animal: animal{
			Name: "le qi",
		},
		age: 2,
	}
	d.run()
	d.rescue()
}

type animal struct {
	Name string
}

type dog struct {
	animal
	age int8
}

func (a animal) run() {
	fmt.Printf("%v is runing! \n", a.Name)
}

func (d dog) rescue() {
	fmt.Printf("%v is rescuing! \n", d.Name)
}
