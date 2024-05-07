package vm

import (
	"encoding/binary"

	"github.com/PhakornKiong/go-vm/opcode"
)

type (
	opFunc func(v *vm) ([]byte, error)
)

type JumpTable [1 << 16]opFunc

func newInstructionSet() JumpTable {
	var jt JumpTable
	jt[opcode.PUSH1] = opPush1
	jt[opcode.PUSH8] = opPush8
	jt[opcode.POP] = opPop
	jt[opcode.ADD] = opAdd
	jt[opcode.SUB] = opSub
	jt[opcode.MUL] = opMul
	jt[opcode.DIV] = opDiv
	jt[opcode.STORE1] = opStore1
	jt[opcode.STORE8] = opStore8
	jt[opcode.LOAD8] = opLoad8
	jt[opcode.RETURN] = opReturn

	for i, entry := range jt {
		if entry == nil {
			jt[i] = opUndefined
		}
	}

	return jt
}

func opUndefined(v *vm) ([]byte, error) {
	return nil, ErrOpUndefined
}

func opPush1(v *vm) ([]byte, error) {
	value := uint64(v.bytecode[v.pc])
	v.stack[v.sp] = value
	v.pc += 1
	v.sp++
	if err := v.checkStackOverflow(); err != nil {
		return nil, err
	}
	return nil, nil
}

func opPush8(v *vm) ([]byte, error) {
	value := binary.BigEndian.Uint64(v.bytecode[v.pc : v.pc+8])
	v.stack[v.sp] = value
	v.pc += 8
	v.sp++
	if err := v.checkStackOverflow(); err != nil {
		return nil, err
	}
	return nil, nil
}

func opPop(v *vm) ([]byte, error) {
	v.sp--
	return nil, nil
}

func opAdd(v *vm) ([]byte, error) {
	if err := v.checkStackUnderflow(uint64(1)); err != nil {
		return nil, err
	}

	a := v.stack[v.sp-1]
	b := v.stack[v.sp-2]
	v.sp -= 2
	v.stack[v.sp] = a + b
	v.sp++
	if err := v.checkStackOverflow(); err != nil {
		return nil, err
	}

	return nil, nil
}

func opSub(v *vm) ([]byte, error) {
	if err := v.checkStackUnderflow(uint64(1)); err != nil {
		return nil, err
	}

	a := v.stack[v.sp-1]
	b := v.stack[v.sp-2]
	v.sp -= 2
	v.stack[v.sp] = a - b
	v.sp++
	if err := v.checkStackOverflow(); err != nil {
		return nil, err
	}

	return nil, nil
}

func opMul(v *vm) ([]byte, error) {
	if err := v.checkStackUnderflow(uint64(1)); err != nil {
		return nil, err
	}
	a := v.stack[v.sp-1]
	b := v.stack[v.sp-2]
	v.sp -= 2
	v.stack[v.sp] = a * b
	v.sp++
	if err := v.checkStackOverflow(); err != nil {
		return nil, err
	}

	return nil, nil
}

func opDiv(v *vm) ([]byte, error) {
	if err := v.checkStackUnderflow(uint64(1)); err != nil {
		return nil, err
	}

	a := v.stack[v.sp-1]
	b := v.stack[v.sp-2]
	v.sp -= 2
	if b == 0 {
		return nil, ErrDivisionByZero
	}
	v.stack[v.sp] = a / b
	v.sp++
	if err := v.checkStackOverflow(); err != nil {
		return nil, err
	}

	return nil, nil
}

func opStore1(v *vm) ([]byte, error) {
	if err := v.checkStackUnderflow(uint64(1)); err != nil {
		return nil, err
	}
	offset := v.stack[v.sp-1]
	value := v.stack[v.sp-2]
	v.sp -= 2
	if err := v.checkMemoryBounds(offset, 1); err != nil {
		return nil, err
	}
	v.memory[offset] = byte(value)

	return nil, nil
}

func opStore8(v *vm) ([]byte, error) {
	if err := v.checkStackUnderflow(uint64(1)); err != nil {
		return nil, err
	}
	offset := v.stack[v.sp-1]
	value := v.stack[v.sp-2]
	v.sp -= 2
	if err := v.checkMemoryBounds(offset, 8); err != nil {
		return nil, err
	}
	binary.BigEndian.PutUint64(v.memory[offset:], value)

	return nil, nil
}

func opLoad8(v *vm) ([]byte, error) {
	if err := v.checkStackUnderflow(uint64(0)); err != nil {
		return nil, err
	}

	offset := v.stack[v.sp-1]
	v.sp--
	if err := v.checkMemoryBounds(offset, 8); err != nil {
		return nil, err
	}
	value := binary.BigEndian.Uint64(v.memory[offset:])
	v.stack[v.sp] = value
	v.sp++
	if err := v.checkStackOverflow(); err != nil {
		return nil, err
	}

	return nil, nil
}

func opReturn(v *vm) ([]byte, error) {
	if err := v.checkStackUnderflow(uint64(1)); err != nil {
		return nil, err
	}
	offset := v.stack[v.sp-1]
	size := v.stack[v.sp-2]
	v.sp -= 2
	if err := v.checkMemoryBounds(offset, 8); err != nil {
		return []byte{}, err
	}
	return v.memory[offset : offset+size], nil
}
