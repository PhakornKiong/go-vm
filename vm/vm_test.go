package vm

import (
	"reflect"
	"testing"

	"github.com/PhakornKiong/go-vm/opcode"
)

func TestVMExecution(t *testing.T) {
	tests := []struct {
		name          string
		bytecode      []byte
		expectedStack []uint64
		err           error
	}{
		{
			name: "Test PUSH1",
			bytecode: []byte{
				byte(opcode.PUSH1), 0x05,
			},
			expectedStack: []uint64{0x05},
			err:           nil,
		},
		{
			name: "Test PUSH8",
			bytecode: []byte{
				byte(opcode.PUSH8), 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05,
			},
			expectedStack: []uint64{0x05},
			err:           nil,
		},
		{
			name: "Test ADD",
			bytecode: []byte{
				byte(opcode.PUSH1), 0x03,
				byte(opcode.PUSH1), 0x05,
				byte(opcode.ADD),
			},
			expectedStack: []uint64{0x08},
			err:           nil,
		},
		{
			name: "Test SUB Positive Result",
			bytecode: []byte{
				byte(opcode.PUSH1), 0x03,
				byte(opcode.PUSH1), 0x05,
				byte(opcode.SUB),
			},
			expectedStack: []uint64{0x02},
			err:           nil,
		},
		{
			name: "Test SUB Negative Result",
			bytecode: []byte{
				byte(opcode.PUSH1), 0x05,
				byte(opcode.PUSH1), 0x03,
				byte(opcode.SUB),
			},
			expectedStack: []uint64{0xFFFFFFFFFFFFFFFE}, // 2's complement of -2
			err:           nil,
		},
		{
			name: "Test MUL",
			bytecode: []byte{
				byte(opcode.PUSH1), 0x03,
				byte(opcode.PUSH1), 0x02,
				byte(opcode.MUL),
			},
			expectedStack: []uint64{0x06},
			err:           nil,
		},
		{
			name: "Test DIV",
			bytecode: []byte{
				byte(opcode.PUSH1), 0x03,
				byte(opcode.PUSH1), 0x06,
				byte(opcode.DIV),
			},
			expectedStack: []uint64{0x02},
			err:           nil,
		},
		{
			name: "Test OPUNDEFINED",
			bytecode: []byte{
				byte(opcode.PUSH1), 0x03,
				byte(opcode.PUSH1), 0x05,
				0xFF, // Undefined opcode
			},
			expectedStack: []uint64{0x03, 0x05}, // Stack should be unchanged due to error
			err:           ErrOpUndefined,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			vm := NewVM()
			vm.LoadBytecode(test.bytecode)

			_, err := vm.Execute()
			if err != test.err {
				t.Errorf("Expected error %v, got %v", test.err, err)
			}
			if !reflect.DeepEqual(vm.stack[:vm.sp], test.expectedStack) {
				t.Errorf("Expected stack %v, got %v", test.expectedStack, vm.stack[:vm.sp])
			}
		})
	}
}

func TestVMStackOverflow(t *testing.T) {
	vm := NewVM()
	vm.sp = uint64(len(vm.stack))

	err := vm.checkStackOverflow()

	if err != ErrStackOverflow {
		t.Errorf("Expected error %v, got %v", ErrStackOverflow, err)
	}
}

func TestVMStackUnderflow(t *testing.T) {
	vm := NewVM()
	vm.sp = uint64(1)

	err := vm.checkStackUnderflow(2)

	if err != ErrStackUnderflow {
		t.Errorf("Expected error %v, got %v", ErrStackUnderflow, err)
	}
}
