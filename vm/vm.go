package vm

import (
	"encoding/binary"
	"fmt"

	"github.com/PhakornKiong/go-vm/opcode"
)

type vm struct {
	stack  []uint64 // stack
	pc     uint64   // program counter
	sp     uint64   // stack pointer
	memory []uint64 // 64-bit addressable memory, each word is 8 byte as well
}

func NewVM() *vm {
	return &vm{
		stack:  make([]uint64, 1<<5),
		pc:     0,
		sp:     0,
		memory: make([]uint64, 1<<5), // Initialize memory with 32-bit addressable space
	}
}

func (v *vm) Execute(bytecode []byte) {
	for v.pc < uint64(len(bytecode)) {
		op := opcode.Opcode(bytecode[v.pc])
		v.pc++
		fmt.Println("Current stack:", v.stack, v.sp)
		switch op {
		case opcode.PUSH:
			if v.pc+8 <= uint64(len(bytecode)) {
				value := binary.LittleEndian.Uint64(bytecode[v.pc : v.pc+8])
				v.stack[v.sp] = value
				v.sp++
				v.pc += 8
			}
		case opcode.POP:
			if v.sp > 0 {
				v.sp--
				v.stack = v.stack[:v.sp]
			}
		case opcode.NEGATE: // TODO Figure out
			if v.sp > 0 {
				v.stack[v.sp-1] = ^v.stack[v.sp-1] + 1 // 2's complement negation
			}
		case opcode.PRINT:
			if v.sp > 0 {
				fmt.Println(v.stack[v.sp-1])
			}
		case opcode.ADD, opcode.MINUS, opcode.MUL, opcode.DIV:
			if v.sp > 1 {
				a := v.stack[v.sp-2]
				b := v.stack[v.sp-1]
				v.sp -= 2
				switch op {
				case opcode.ADD:
					v.stack[v.sp] = a + b
				case opcode.MINUS:
					v.stack[v.sp] = a - b
				case opcode.MUL:
					v.stack[v.sp] = a * b
				case opcode.DIV:
					if b != 0 {
						v.stack[v.sp] = a / b
					} else {
						fmt.Println("Error: Division by zero")
					}
				}
				v.sp++
			}
		case opcode.STORE:
			if v.sp > 1 {
				memoryOffset := v.stack[v.sp-2]
				valueToStore := v.stack[v.sp-1]
				v.sp -= 2
				if memoryOffset < uint64(len(v.memory)) {
					v.memory[memoryOffset] = valueToStore
				} else {
					fmt.Println("Error: Memory access out of bounds")
				}
			}
		case opcode.LOAD:
			if v.sp > 0 {
				memoryOffset := v.stack[v.sp-1]
				v.sp--
				if memoryOffset < uint64(len(v.memory)) {
					v.stack[v.sp] = v.memory[memoryOffset]
					v.sp++
				} else {
					fmt.Println("Error: Memory access out of bounds")
				}
			}
		}
	}
}
