package opcode

import "fmt"

type Opcode byte

const (
	// Unary opcodes
	// POP is used to retrieve a value from the stack.
	POP Opcode = iota
	// PUSH is used to place a value onto the stack.
	PUSH
	// NEGATE is used to negate a single operand.
	NEGATE
	// PRINT is used to print the top value of the stack as uint64
	PRINT
	// PRINT_INT64 is used to print the top value of the stack as int64
	PRINT_INT64
	// LOAD is used to load a value from memory into the stack at a specified address.
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
	// STORE is used to store a value from the stack into memory at a specified address.
	STORE
)

func (op Opcode) IsUnary() bool {
	switch op {
	case POP, PUSH, NEGATE:
		return true
	default:
		return false
	}
}

func (op Opcode) IsBinary() bool {
	switch op {
	case ADD, SUB, MUL, DIV:
		return true
	default:
		return false
	}
}

func FromString(opStr string) (Opcode, error) {
	switch opStr {
	case "POP":
		return POP, nil
	case "PUSH":
		return PUSH, nil
	case "NEGATE":
		return NEGATE, nil
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
	case NEGATE:
		return "NEGATE"
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
