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

func (c *Compiler) Compile(input string) error {
	// Each line is a command
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
			return err
		}

		switch op {
		case opcode.PUSH:
			if len(parts) > 1 {
				var buf []byte

				value, err := strconv.ParseUint(parts[1], 10, 64)
				if err != nil {
					fmt.Println("Error parsing PUSH operand as uint64:", err)
					return err
				}

				buf = make([]byte, 8)
				binary.BigEndian.PutUint64(buf, value)

				c.bytecode = append(c.bytecode, byte(opcode.PUSH))
				c.bytecode = append(c.bytecode, buf...)
			}
		case opcode.ADD:
			c.bytecode = append(c.bytecode, byte(opcode.ADD))
		case opcode.SUB:
			c.bytecode = append(c.bytecode, byte(opcode.SUB))
		case opcode.PRINT:
			c.bytecode = append(c.bytecode, byte(opcode.PRINT))
		case opcode.PRINT_INT64:
			c.bytecode = append(c.bytecode, byte(opcode.PRINT_INT64))
		case opcode.STORE:
			c.bytecode = append(c.bytecode, byte(opcode.STORE))
		case opcode.LOAD:
			c.bytecode = append(c.bytecode, byte(opcode.LOAD))
		default:
			return fmt.Errorf("opcode not found")
		}

	}

	return nil
}

func (c *Compiler) Output() []byte {
	return (c.bytecode)
}
