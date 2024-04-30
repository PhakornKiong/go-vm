package main

import "github.com/PhakornKiong/go-vm/vm"

func main() {
	a := vm.NewVM()

	a.Execute(
		[]vm.Instruction{
			// {Opcode: vm.PUSH, Operands: []int64{1}},
			// {Opcode: vm.PUSH, Operands: []int64{1}},
			// {Opcode: vm.PRINT},
			// {Opcode: vm.ADD},
			// {Opcode: vm.PRINT},
			// {Opcode: vm.NEGATE},
			// {Opcode: vm.PRINT},
			{Opcode: vm.PUSH, Operands: []int64{9}},
			{Opcode: vm.PUSH, Operands: []int64{20}},
			{Opcode: vm.STORE},
			{Opcode: vm.PUSH, Operands: []int64{9}},
			{Opcode: vm.LOAD},
			{Opcode: vm.PRINT},
		},
	)
}
