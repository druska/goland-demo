package main

import (
	"fmt"

	"github.com/druska/goland-demo/modulea"
)

func main() {
	fmt.Println("Hello from Module B")
	modulea.ModuleAFunction()
}
