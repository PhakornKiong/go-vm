package main

import (
	"github.com/PhakornKiong/go-vm/compiler"
	"github.com/PhakornKiong/go-vm/vm"
)

func main() {
	a := vm.NewVM()

	compiler := compiler.NewCompiler()
	compiler.Compile(`
		PUSH 5000
		PUSH 6000
		PRINT
		SUB
		PRINT_INT64
		PRINT
	`)

	a.Execute(compiler.Output())

}
