package vm

import (
	"fmt"
)

var (
	ErrMemoryOutOfBounds = fmt.Errorf("memory access out of bounds")
	ErrDivisionByZero    = fmt.Errorf("division by zero")
	ErrOpUndefined       = fmt.Errorf("undefined opcode")

	ErrStackUnderflow = fmt.Errorf("stack underflow")
	ErrStackOverflow  = fmt.Errorf("stack overflow")
)

type vm struct {
	stack     []uint64 // stack
	pc        uint64   // program counter
	sp        uint64   // stack pointer
	bytecode  []byte   // compiled bytecode
	memory    []byte   // memory
	jumpTable JumpTable
}

func NewVM() *vm {
	return &vm{
		stack:     make([]uint64, 32), // stack depth of 32
		pc:        0,
		sp:        0,
		memory:    make([]byte, 1<<16), // 64 KB of byte addressable memory
		jumpTable: newInstructionSet(),
	}
}

func (v *vm) LoadBytecode(bytecode []byte) {
	v.bytecode = bytecode
}

func (v *vm) Execute() ([]byte, error) {
	for v.pc < uint64(len(v.bytecode)) {
		op := v.jumpTable[v.bytecode[v.pc]]
		v.pc++
		res, err := op(v)

		if err != nil {
			return nil, err
		}

		if res != nil {
			return res, nil
		}
	}
	return []byte{}, nil
}

func (v *vm) checkMemoryBounds(address uint64, length uint64) error {
	if address+length > uint64(len(v.memory)) {
		return ErrMemoryOutOfBounds
	}
	return nil
}

func (v *vm) checkStackUnderflow(required uint64) error {
	if v.sp < required {
		return ErrStackUnderflow
	}
	return nil
}

func (v *vm) checkStackOverflow() error {
	if v.sp >= uint64(len(v.stack)) {
		return ErrStackOverflow
	}
	return nil
}
