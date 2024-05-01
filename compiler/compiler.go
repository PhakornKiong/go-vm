package compiler

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"

	"github.com/PhakornKiong/go-vm/opcode"
)

type Compiler struct {
	bytecode []byte
}

func NewCompiler() *Compiler {
	return &Compiler{
		bytecode: make([]byte, 0),
	}
}

func (c *Compiler) Compile(input string) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		// Assume that first word is always the opcode
		op, err := opcode.FromString(parts[0])
		if err != nil {
			fmt.Println("Error converting string to opcode:", err)
			continue
		}
		switch op {
		case opcode.PUSH:
			if len(parts) > 1 {
				value, err := strconv.ParseUint(parts[1], 10, 64)
				if err != nil {
					fmt.Println("Error parsing PUSH operand:", err)
					continue
				}
				c.bytecode = append(c.bytecode, byte(opcode.PUSH))
				// Convert uint64 to bytes and append
				buf := make([]byte, 8)
				binary.LittleEndian.PutUint64(buf, value)
				c.bytecode = append(c.bytecode, buf...)
			}
		case opcode.ADD:
			c.bytecode = append(c.bytecode, byte(opcode.ADD))
		case opcode.PRINT:
			c.bytecode = append(c.bytecode, byte(opcode.PRINT))
		}
	}
}

func (c *Compiler) Output() []byte {
	return (c.bytecode)
}
