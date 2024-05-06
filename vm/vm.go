package vm

import (
	"encoding/binary"
	"fmt"

	"github.com/PhakornKiong/go-vm/opcode"
)

const (
	ErrMemoryOutOfBounds = "memory access out of bounds"
)

func (v *vm) checkMemoryBounds(address uint64, length uint64) error {
	if address+length > uint64(len(v.memory)) {
		return fmt.Errorf(ErrMemoryOutOfBounds)
	}
	return nil
}

type vm struct {
	stack  []uint64 // stack
	pc     uint64   // program counter
	sp     uint64   // stack pointer
	memory []byte   // memory
}

func NewVM() *vm {
	return &vm{
		stack:  make([]uint64, 32), // stack depth of 32
		pc:     0,
		sp:     0,
		memory: make([]byte, 1<<8), // 4 GB of byte addressable memory
	}
}

func (v *vm) Execute(bytecode []byte) ([]byte, error) {
	for v.pc < uint64(len(bytecode)) {
		op := opcode.Opcode(bytecode[v.pc])
		v.pc++

		// fmt.Println("Current stack state:", v.stack)
		// fmt.Println("Current memory state:", v.memory)

		switch op {
		case opcode.PUSH1:
			value := uint64(bytecode[v.pc])
			v.stack[v.sp] = value
			v.sp++
			v.pc += 1
		case opcode.PUSH8:
			value := binary.BigEndian.Uint64(bytecode[v.pc : v.pc+8])
			v.stack[v.sp] = value
			v.sp++
			v.pc += 8
		case opcode.POP:
			v.sp--
		case opcode.ADD, opcode.SUB, opcode.MUL, opcode.DIV:
			if v.sp > 1 {
				a := v.stack[v.sp-1]
				b := v.stack[v.sp-2]
				v.sp -= 2
				switch op {
				case opcode.ADD:
					v.stack[v.sp] = a + b
				case opcode.SUB:
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
		case opcode.STORE1:
			if v.sp > 1 {
				offset := v.stack[v.sp-1]
				value := v.stack[v.sp-2]
				v.sp -= 2
				if err := v.checkMemoryBounds(offset, 1); err != nil {
					return []byte{}, err
				}
				buf := make([]byte, 1)
				buf[0] = byte(value) // Take only the least significant byte of the value
				copy(v.memory[offset:], buf)
			}
		case opcode.STORE8:
			if v.sp > 1 {
				offset := v.stack[v.sp-1]
				value := v.stack[v.sp-2]
				v.sp -= 2
				if err := v.checkMemoryBounds(offset, 8); err != nil {
					return []byte{}, err
				}
				buf := make([]byte, 8)
				binary.BigEndian.PutUint64(buf, value)
				copy(v.memory[offset:], buf)
			}
		case opcode.LOAD8:
			if v.sp > 1 {
				offset := v.stack[v.sp-1]
				v.sp--
				if err := v.checkMemoryBounds(offset, 8); err != nil {
					return []byte{}, err
				}
				value := binary.BigEndian.Uint64(v.memory[offset:])
				v.stack[v.sp] = value
				v.sp++
			}
		case opcode.RETURN:
			if v.sp > 1 {
				offset := v.stack[v.sp-1]
				size := v.stack[v.sp-2]
				v.sp -= 2
				if err := v.checkMemoryBounds(offset, 8); err != nil {
					return []byte{}, err
				}
				return v.memory[offset : offset+size], nil
			}
		}
	}
	return []byte{}, nil
}
