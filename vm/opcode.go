package vm

type Opcode byte

const (
	// Unary opcodes
	// POP is used to retrieve a value from the stack.
	POP Opcode = iota
	// PUSH is used to place a value onto the stack.
	PUSH
	// NEGATE is used to negate a single operand.
	NEGATE
	// PRINT is used to print the top value of the stack.
	PRINT
	// LOAD is used to load a value from memory into the stack at a specified address.
	LOAD

	// Binary opcodes
	// ADD is used to add two operands.
	ADD
	// MINUS is used to subtract the second operand from the first operand.
	MINUS
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
	case ADD, MINUS, MUL, DIV:
		return true
	default:
		return false
	}
}
