package main

import "fmt"

type Human struct {
	name string
	ability string
	age int
	sameVar int
}

func (h *Human) greeting(){
	fmt.Printf("I am %s, and i am able to %s. I have lived for %d long years\n", h.name, h.ability, h.age)
}

func (h *Human) sameMethod(){
	fmt.Println("sameMethod call from Human")
}

type Action struct {
	Human
	sameVar int
}

func (a *Action) doit(){
	fmt.Printf("%s, uses his ability %s, for sake of human kind\n", a.name, a.ability)
}

func (a *Action) sameMethod(){
	fmt.Println("sameMethod call from Action")
}


func main(){
	fmt.Printf("--------------Creating Bob struct, calling greeting() on him, printing his fields\n")
	h1 := Human{"Bob", "sleep", 10, -1}
	h1.greeting()
	fmt.Printf("Name: %s, abillity: %s, age: %d, sameVar: %d\n", h1.name, h1.ability, h1.age, h1.sameVar)
	fmt.Printf("---\n\n")

	fmt.Printf("--------------Creating Action struct with Bob(Human) embedded in it, calling doit(), greeting() on him, printing his fields\n")
	a1 := Action{h1, 1}
	a1.doit()
	a1.greeting()
	fmt.Printf("Name: %s, abillity: %s, age: %d, sameVar: %d\n", a1.name, a1.ability, a1.age, a1.sameVar)
	fmt.Printf("---\n\n")

	fmt.Printf("--------------Both Action and Human struct has method sameMethod(), calling it on Action will dominate over embedded struct\n")
	a1.sameMethod()
	fmt.Printf("--------------Same applie to variable\n")
	fmt.Printf("Action struct .sameVar filed: %d, Human struct .sameVar field %d\n", h1.sameVar, a1.sameVar)
	fmt.Printf("---\n\n")

	fmt.Printf("Composition is easier, than creating a stuct field, cause in later case we will have to create wrapper for each method of stuct\n")
	fmt.Printf("Also, we can embed struct pointer rather than value. Then, our inherited struct will modify embeded struct\n")
}

