# A Simple Go Virtual Machine (VM)

## Overview

This project implements a simple virtual machine (VM) in Go, capable of executing a set of predefined operations encoded as bytecode. The VM operates on a byte-addressable memory model and supports basic arithmetic and memory storage.

## Design

The VM is designed as a stack-based system utilizing a 64-bit architecture, where each word is precisely 8 bytes. This setup leverages `uint64` for all stack operations and memory addressing within a 64 KB byte-addressable memory space. It employs a big endian format for byte ordering in memory operations.

### Stack Mechanism

The stack in this VM operates on a Last In, First Out (LIFO) principle, where data is pushed onto and popped from the top of the stack.

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

## Example

TODO
