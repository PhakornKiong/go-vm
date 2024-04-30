package vm

import "fmt"

type vm struct {
	stack  []int64 // stack
	sp     int64   // stack pointer
	memory []int64 // 64-bit addressable memory, each word is 8 byte as well
}

type Instruction struct {
	Opcode   Opcode
	Operands []int64
}

func NewVM() *vm {
	return &vm{
		stack:  make([]int64, 1<<5),
		sp:     0,
		memory: make([]int64, 1<<5), // Initialize memory with 32-bit addressable space
	}
}

func (v *vm) Execute(instructions []Instruction) {
	for _, instr := range instructions {
		fmt.Println("Current stack:", v.stack, v.sp)
		switch instr.Opcode {
		case PUSH:
			if len(instr.Operands) > 0 {
				v.stack[v.sp] = instr.Operands[0]
				v.sp++
			}
		case POP:
			if v.sp > 0 {
				v.sp--
				v.stack = v.stack[:v.sp]
			}
		case NEGATE:
			if v.sp > 0 {
				v.stack[v.sp-1] = -v.stack[v.sp-1]
			}
		case PRINT:
			if v.sp > 0 {
				fmt.Println(v.stack[v.sp-1])
			}
		case ADD, MINUS, MUL, DIV:
			if v.sp > 1 {
				a := v.stack[v.sp-2]
				b := v.stack[v.sp-1]
				v.sp -= 2
				switch instr.Opcode {
				case ADD:
					v.stack[v.sp] = a + b
				case MINUS:
					v.stack[v.sp] = a - b
				case MUL:
					v.stack[v.sp] = a * b
				case DIV:
					if b != 0 {
						v.stack[v.sp] = a / b
					} else {
						fmt.Println("Error: Division by zero")
					}
				}
				v.sp++
			}
		case STORE:
			if v.sp > 1 {
				memoryOffset := v.stack[v.sp-2]
				valueToStore := v.stack[v.sp-1]
				v.sp -= 2
				// assume memory offset within possible range

				v.memory[memoryOffset] = valueToStore
			}
		case LOAD:
			if v.sp > 0 {
				memoryOffset := v.stack[v.sp-1]
				v.sp--
				// assume memory offset within possible range
				fmt.Println("Memory Content:", v.memory, v.sp, v.memory[memoryOffset])
				if memoryOffset >= 0 && memoryOffset < int64(len(v.memory)) {
					v.stack[v.sp] = v.memory[memoryOffset]
					v.sp++
					fmt.Println("Stack Pointer (v.sp):", v.sp)

				} else {
					fmt.Println("Error: Memory access out of bounds")
				}
			}
		}

	}

}
