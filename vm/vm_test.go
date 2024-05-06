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
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			vm := NewVM()
			_, err := vm.Execute(test.bytecode)
			if err != test.err {
				t.Errorf("Expected error %v, got %v", test.err, err)
			}
			if !reflect.DeepEqual(vm.stack[:vm.sp], test.expectedStack) {
				t.Errorf("Expected stack %v, got %v", test.expectedStack, vm.stack[:vm.sp])
			}
		})
	}
}
