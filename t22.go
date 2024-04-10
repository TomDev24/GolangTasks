package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b,
// значение которых > 2^20.

func main() {
	a := big.NewInt(1 << 30)
	b := big.NewInt(1 << 40)
	var c big.Int
	var res big.Int
	c.SetString("1267650600228229401496703205376", 10)

	fmt.Printf("a: %d, b: %d, c: %s\n\n", a, b, c.String())

	fmt.Printf("%d + %d = %d\n", a, b, res.Add(a,b))
	fmt.Printf("%s - %d = %d\n\n", res.String(), b, res.Sub(&res,b))
	fmt.Printf("%d * %d = %d\n", a, b, res.Mul(a,b))
	fmt.Printf("%s / %d = %d\n", res.String(), b, res.Div(&res,b))
}
