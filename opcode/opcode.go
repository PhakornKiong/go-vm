package opcode

import "fmt"

type Opcode byte

var ErrUnknownOpcode = fmt.Errorf("unknown opcode")

const (
	// Unary opcodes
	// POP is used to retrieve a value from the stack.
	POP Opcode = iota
	// PUSH1 is used to place a 1 byte onto the stack.
	PUSH1
	// PUSH8 is used to place a 8 byte (a word) onto the stack.
	PUSH8
	// LOAD is used to load 8 byte (a word) from memory into the stack at a specified offset.
	LOAD8

	// Binary opcodes
	// ADD is used to add two operands.
	ADD
	// SUB is used to subtract the second operand from the first operand.
	SUB
	// MUL is used to multiply two operands.
	MUL
	// DIV is used to divide the first operand by the second operand.
	DIV
	// STORE1 is used to store 1 byte into memory at a specified offset.
	STORE1
	// STORE8 is used to store 8 byte (a word) into memory at a specified offset.
	STORE8
	// RETURN is exit the execution and return the offset from memory
	RETURN
)

func FromString(opStr string) (Opcode, error) {
	switch opStr {
	case "POP":
		return POP, nil
	case "PUSH1":
		return PUSH1, nil
	case "PUSH8":
		return PUSH8, nil
	case "LOAD8":
		return LOAD8, nil
	case "ADD":
		return ADD, nil
	case "SUB":
		return SUB, nil
	case "MUL":
		return MUL, nil
	case "DIV":
		return DIV, nil
	case "STORE1":
		return STORE1, nil
	case "STORE8":
		return STORE8, nil
	case "RETURN":
		return RETURN, nil
	default:
		return 0, ErrUnknownOpcode
	}
}
