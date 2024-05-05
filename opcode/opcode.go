package opcode

import "fmt"

type Opcode byte

const (
	// Unary opcodes
	// POP is used to retrieve a value from the stack.
	POP Opcode = iota
	// PUSH is used to place a value onto the stack.
	PUSH
	// PRINT is used to print the top value of the stack as uint64
	PRINT
	// PRINT_INT64 is used to print the top value of the stack as int64
	PRINT_INT64
	// LOAD is used to load an entire word from memory into the stack at a specified address.
	LOAD

	// Binary opcodes
	// ADD is used to add two operands.
	ADD
	// SUB is used to subtract the second operand from the first operand.
	SUB
	// MUL is used to multiply two operands.
	MUL
	// DIV is used to divide the first operand by the second operand.
	DIV
	// STORE is used to store an entire word from the stack into memory at a specified address.
	STORE
)

func FromString(opStr string) (Opcode, error) {
	switch opStr {
	case "POP":
		return POP, nil
	case "PUSH":
		return PUSH, nil
	case "PRINT":
		return PRINT, nil
	case "PRINT_INT64":
		return PRINT_INT64, nil
	case "LOAD":
		return LOAD, nil
	case "ADD":
		return ADD, nil
	case "SUB":
		return SUB, nil
	case "MUL":
		return MUL, nil
	case "DIV":
		return DIV, nil
	case "STORE":
		return STORE, nil
	default:
		return 0, fmt.Errorf("unknown opcode: %s", opStr)
	}
}

func (op Opcode) String() string {
	switch op {
	case POP:
		return "POP"
	case PUSH:
		return "PUSH"
	case PRINT:
		return "PRINT"
	case PRINT_INT64:
		return "PRINT_INT64"
	case LOAD:
		return "LOAD"
	case ADD:
		return "ADD"
	case SUB:
		return "SUB"
	case MUL:
		return "MUL"
	case DIV:
		return "DIV"
	case STORE:
		return "STORE"
	default:
		return "UNKNOWN"
	}
}
