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
		PUSH1 0xff
		PUSH1 0x00
		STORE1
		PUSH1 0x01
		PUSH1 0x00
		RETURN
	`)
	output := compiler.Output()
	fmt.Print("Bytecode Output: ")
	for _, b := range output {
		fmt.Printf("%02x ", b)
	}
	fmt.Println()

	res, _ := a.Execute(compiler.Output())
	fmt.Println(res)

}
