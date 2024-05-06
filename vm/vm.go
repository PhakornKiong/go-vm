package vm

import (
	"encoding/binary"
	"fmt"

	"github.com/PhakornKiong/go-vm/opcode"
)

var (
	ErrMemoryOutOfBounds = fmt.Errorf("memory access out of bounds")
	ErrDivisionByZero    = fmt.Errorf("division by zero")
)

func (v *vm) checkMemoryBounds(address uint64, length uint64) error {
	if address+length > uint64(len(v.memory)) {
		return ErrMemoryOutOfBounds
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
		memory: make([]byte, 1<<16), // 64 KB of byte addressable memory
	}
}

func (v *vm) Execute(bytecode []byte) ([]byte, error) {
	for v.pc < uint64(len(bytecode)) {
		op := opcode.Opcode(bytecode[v.pc])
		v.pc++

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
		case opcode.ADD:
			if v.sp > 1 {
				a := v.stack[v.sp-1]
				b := v.stack[v.sp-2]
				v.sp -= 2
				v.stack[v.sp] = a + b
				v.sp++
			}
		case opcode.SUB:
			if v.sp > 1 {
				a := v.stack[v.sp-1]
				b := v.stack[v.sp-2]
				v.sp -= 2
				v.stack[v.sp] = a - b
				v.sp++
			}
		case opcode.MUL:
			if v.sp > 1 {
				a := v.stack[v.sp-1]
				b := v.stack[v.sp-2]
				v.sp -= 2
				v.stack[v.sp] = a * b
				v.sp++
			}
		case opcode.DIV:
			if v.sp > 1 {
				a := v.stack[v.sp-1]
				b := v.stack[v.sp-2]
				v.sp -= 2
				if b == 0 {
					return nil, ErrDivisionByZero
				}

				v.stack[v.sp] = a / b
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
