package main

import "fmt"

type Human struct {
	name string
	ability string
	age int
}

func (h *Human) greeting(){
	fmt.Printf("I am %s, and i am able to %s. I have lived for %d long years\n", h.name, h.ability, h.age)
}

type Action struct {
	Human
}

func (a *Action) doit(){
	fmt.Printf("%s, uses his ability %s, for sake of human kind\n", a.name, a.ability)
}


func main(){
	h1 := Human{"Bob", "sleep", 10}
	h1.greeting()
	fmt.Println(h1.name, h1.ability, h1.age)

	fmt.Println()
	a1 := Action{h1}
	a1.doit()
	a1.greeting()
	fmt.Println(a1.name, a1.ability, a1.age)
}
