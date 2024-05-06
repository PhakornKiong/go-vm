package opcode

import "fmt"

type Opcode byte

var ErrUnknownOpcode = fmt.Errorf("unknown opcode")

const (
	// POP is used to discard the top value from the stack and decrement the stack pointer.
	POP Opcode = iota
	// PUSH1 is used to place a 1 byte data onto the stack.
	PUSH1
	// PUSH8 is used to place a 8 byte (a word) data onto the stack.
	PUSH8
	// LOAD is used to load 8 byte (a word) data from memory into the stack at a specified offset using the top value of the stack
	LOAD8

	// ADD is used to add two operands.
	ADD
	// SUB subtracts the top value of the stack from the second top value of the stack.
	SUB
	// MUL multiplies the top two values on the stack.
	MUL
	// DIV divides the top value of the stack by the second top value of the stack, errors on division by zero.
	DIV
	// STORE1 stores a 1-byte value into memory at a specified offset, using the top value of the stack as the offset and the second top value as the data to store
	STORE1
	// STORE8 stores an 8-byte value (a word) into memory at a specified offset, using the top value of the stack as the offset and the second top value as the data to store
	STORE8
	// RETURN is used to exit the execution and return a block of memory.
	// It pops the size and offset from the stack and returns the specified block from memory.
	RETURN
)

var opcodeMap = map[string]Opcode{
	"POP":    POP,
	"PUSH1":  PUSH1,
	"PUSH8":  PUSH8,
	"LOAD8":  LOAD8,
	"ADD":    ADD,
	"SUB":    SUB,
	"MUL":    MUL,
	"DIV":    DIV,
	"STORE1": STORE1,
	"STORE8": STORE8,
	"RETURN": RETURN,
}

func FromString(opStr string) (Opcode, error) {
	if op, exists := opcodeMap[opStr]; exists {
		return op, nil
	}
	return 0, ErrUnknownOpcode
}
