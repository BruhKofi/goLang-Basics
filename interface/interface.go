package main

import (
	"fmt"
)

func main() {
	p := new(person)
	fmt.Println(p.talk())
	fmt.Println(p.walk())

	o := new(policeOfficer)
	fmt.Println(o.talk())
	fmt.Println(o.walk())
	o.badgeNumber = 1000
	fmt.Println(o.badgeNumber)
	fmt.Println(o.run())
	fmt.Println(run(12))
}

type person struct {
	firstname string
	lastname  string
}

type policeOfficer struct {
	badgeNumber int
	precint     string
}

type behaviours interface {
	talk() string
	walk() int
	run() int
}

func (p *person) talk() string {
	return "hi there"
}

func (p *person) walk() int {
	return 10
}

//officer implementation

func (o *policeOfficer) talk() string {
	return "have a good day"
}

func (o *policeOfficer) walk() int {
	return 20
}

//funct [param list][interface func name][interface func return type]
func (o *policeOfficer) run() int {
	return 50
}

//regular function
//func [name][param][param list][return type]
func run(s int) int {
	return s
}
