package main

import (
	"github.com/PhakornKiong/go-vm/compiler"
	"github.com/PhakornKiong/go-vm/vm"
)

func main() {
	a := vm.NewVM()

	compiler := compiler.NewCompiler()
	compiler.Compile(`
		PUSH 1
		PUSH 1
		PRINT
		ADD
		PRINT
		NEGATE
		PRINT
	`)

	a.Execute(compiler.Output())

}
