# A Simple Go Virtual Machine (VM)

## Overview

This project implements a simple virtual machine (VM) in Go, capable of executing a set of predefined operations encoded as bytecode. The VM operates on a byte-addressable memory model and supports basic arithmetic and memory storage.

## Design

The VM is designed as a stack-based system utilizing a 64-bit architecture, where each word is precisely 8 bytes. This setup leverages `uint64` for all stack operations and memory addressing within a 64 KB byte-addressable memory space. It employs a big endian format for byte ordering in memory operations.

### Stack Mechanism

The stack in this VM operates on a Last In, First Out (LIFO) principle, where data is pushed onto and popped from the top of the stack.

The **stack pointer (SP)** in the VM is a crucial component that helps in managing the stack's state during the execution of operations. The stack pointer points to the next free position on the stack where data can be pushed. When data is pushed onto the stack, it is placed at the position indicated by the stack pointer, and then the stack pointer is incremented to point to the next free position.

Conversely, when data is popped from the stack, the stack pointer is decremented first, and then the data at the new stack pointer position is retrieved. This ensures that the last pushed item is the first to be popped, adhering to the Last In, First Out (LIFO) principle.

Here's a brief overview of how the stack pointer changes with various operations:

- **PUSH Operations**: When a value is pushed onto the stack, it is stored at the current position of the stack pointer. After storing the value, the stack pointer is incremented by one position regardless of the size of the data pushed (`PUSH1` or `PUSH8`).

- **POP Operation**: Before a value is popped from the stack, the stack pointer is decremented by the size of the data to be popped (1 byte for `POP` when it retrieves the last pushed byte-sized data). The value at the new stack pointer position is then retrieved.

- **Arithmetic Operations (ADD, SUB, MUL, DIV)**: These operations typically involve popping the top two values from the stack, performing the arithmetic operation, and then pushing the result back onto the stack. The stack pointer is adjusted accordingly after each pop and push operation.

The management of the stack pointer is automatic and internal to the VM's operation, ensuring that the stack's integrity is maintained throughout the execution of the bytecode.

## Opcodes

The VM recognizes the following opcodes:

- **PUSH1**: Pushes a 1-byte value onto the stack.
- **PUSH8**: Pushes an 8-byte value (a word) onto the stack.
- **POP**: Pops the top value from the stack.
- **ADD**: Adds the top two values on the stack.
- **SUB**: Subtracts the second top value from the top value on the stack.
- **MUL**: Multiplies the top two values on the stack.
- **DIV**: Divides the top value by the second top value on the stack, errors on division by zero.
- **STORE1**: Stores a 1-byte value into memory at a specified offset.
- **STORE8**: Stores an 8-byte value (a word) into memory at a specified offset.
- **LOAD8**: Loads an 8-byte value (a word) from memory into the stack at a specified offset.
- **RETURN**: Returns a block of memory and halts execution.

Here is an example demonstrating the use of the `PUSH` and `ADD` opcodes in our VM:

### Example: Using PUSH and ADD Opcodes

This example demonstrates the use of `PUSH1` and `ADD` opcodes and visualizes the stack state at each step:

1. **PUSH1 0x03**: Pushes the hexadecimal value `0x03` (which is `3` in decimal) onto the stack.

   - Stack state after operation: `[3]`

2. **PUSH1 0x05**: Pushes the hexadecimal value `0x05` (which is `5` in decimal) onto the stack.

   - Stack state after operation: `[3, 5]`

3. **ADD**: Pops the top two values from the stack (`3` and `5`), adds them together resulting in `8`, and pushes the result back onto the stack.
   - Stack state after operation: `[8]`

The sequence of these opcodes in bytecode and their execution by the VM results in the final top value of the stack being `8`. This visualization helps in understanding how values are manipulated on the stack during the execution of opcodes.

## Usage

To use the VM, compile bytecode using the provided compiler which translates a set of instructions into bytecode that the VM can execute. The VM then executes the bytecode and manipulates its internal state according to the instructions encoded in the bytecode.

To run the virtual machine (VM) and execute bytecode, you can use the `main.go` file with specific command-line flags. Here's how you can use it:

1. **Input File**: Use the `-i` or `--input` flag to specify the input file containing the opcode instructions. This file should be a plain text file with the opcodes that you want the VM to execute.

   Example:

   ```
   go run main.go --input examples/addition
   ```

2. **Return Type**: Use the `-t` or `--type` flag to specify the type of the return value. The supported types are `int64`, `uint64`, and `string`. If no type is specified, the result will be output as a byte array.

   Example:

   ```
   go run main.go --input examples/subtraction --type int64
   ```
