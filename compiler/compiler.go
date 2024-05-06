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
		case opcode.PUSH1:
			if len(parts) > 1 {
				var buf []byte

				value, err := strconv.ParseUint(parts[1][2:], 16, 8) // Parse hexadecimal, skipping the '0x' prefix
				if err != nil {
					fmt.Println("Error parsing PUSH1 operand as uint8 in hexadecimal:", err)
					return err
				}

				buf = make([]byte, 1)
				buf[0] = byte(value)

				c.bytecode = append(c.bytecode, byte(opcode.PUSH1))
				c.bytecode = append(c.bytecode, buf...)
			}
		case opcode.PUSH8:
			if len(parts) > 1 {
				var buf []byte

				value, err := strconv.ParseUint(parts[1][2:], 16, 64) // Parse hexadecimal, skipping the '0x' prefix
				if err != nil {
					fmt.Println("Error parsing PUSH8 operand as uint64 in hexadecimal:", err)
					return err
				}

				buf = make([]byte, 8)
				binary.BigEndian.PutUint64(buf, value)

				c.bytecode = append(c.bytecode, byte(opcode.PUSH8))
				c.bytecode = append(c.bytecode, buf...)
			}
		case opcode.ADD:
			c.bytecode = append(c.bytecode, byte(opcode.ADD))
		case opcode.SUB:
			c.bytecode = append(c.bytecode, byte(opcode.SUB))
		case opcode.STORE1:
			c.bytecode = append(c.bytecode, byte(opcode.STORE1))
		case opcode.STORE8:
			c.bytecode = append(c.bytecode, byte(opcode.STORE8))
		case opcode.LOAD8:
			c.bytecode = append(c.bytecode, byte(opcode.LOAD8))
		case opcode.RETURN:
			c.bytecode = append(c.bytecode, byte(opcode.RETURN))
		default:
			return fmt.Errorf("opcode not found")
		}

	}

	return nil
}

func (c *Compiler) Output() []byte {
	return (c.bytecode)
}
