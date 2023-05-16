package go_context

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b"))
	fmt.Println(contextF.Value("a"))

	fmt.Println(contextE.Value("e"))
	fmt.Println(contextE.Value("b"))
	fmt.Println(contextE.Value("a"))

	fmt.Println(contextD.Value("d"))
	fmt.Println(contextD.Value("b"))
	fmt.Println(contextD.Value("a"))

	fmt.Println(contextC.Value("c"))
	fmt.Println(contextC.Value("a"))

	fmt.Println(contextB.Value("b"))
	fmt.Println(contextB.Value("a"))

	fmt.Println(contextA.Value("a"))
}