package main

import (
	"fmt"

	"github.com/PhakornKiong/go-vm/compiler"
	"github.com/PhakornKiong/go-vm/vm"
)

func main() {
	a := vm.NewVM()

	compiler := compiler.NewCompiler()
	compiler.Compile(`
		PUSH 7000
		PUSH 6000
		PRINT
		SUB
		PRINT_INT64
		PRINT
		PUSH 50
		PUSH 0
		STORE
		PUSH 0
		LOAD
		PRINT
		PUSH 16
		PUSH 0
		RETURN
	`)

	res, _ := a.Execute(compiler.Output())
	fmt.Println(res)

}
