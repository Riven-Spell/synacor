package tests

import (
	"synacor/interpreter"
	"synacor/interpreter/opcodes"
	"testing"
)

func loadEfficient(system *interpreter.System, toLoad []uint16) {
	for k, v := range toLoad {
		system.Memory[k] = v
	}
}

//Add 4 to value in register 1 (0), store in register 0
func TestAdd(t *testing.T) {
	system := interpreter.NewSystem()
	loadEfficient(system, []uint16{opcodes.ADD, 32768, 32769, 4})

	if err := system.Step(); err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if system.Registers[0] != 4 {
		t.Errorf("Expected system.Registers[0] to be 4, got %d\n", system.Registers[0])
		t.Fail()
	}
}

//Multiply 4 by 5, store in register 0, result 20
func TestMul(t *testing.T) {
	system := interpreter.NewSystem()
	loadEfficient(system, []uint16{opcodes.MULT, 32768, 4, 5})

	if err := system.Step(); err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if system.Registers[0] != 20 {
		t.Errorf("Expected system.Registers[0] to be 20, got %d\n", system.Registers[0])
		t.Fail()
	}
}

//modulo 21 by 4, store in register 0, result 1
func TestModulo(t *testing.T) {
	system := interpreter.NewSystem()
	loadEfficient(system, []uint16{opcodes.MOD, 32768, 21, 4})

	if err := system.Step(); err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if system.Registers[0] != 1 {
		t.Errorf("Expected system.Registers[0] to be 1, got %d\n", system.Registers[0])
		t.Fail()
	}
}

//Halt the system.
func TestHalt(t *testing.T) {
	system := interpreter.NewSystem()
	system.Memory[0] = opcodes.HALT

	if err := system.Step(); err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if system.Halted != true {
		t.Error("Expected system to be halted.")
		t.Fail()
	}
}

//Jump forward in memory.
func TestJump(t *testing.T) {
	system := interpreter.NewSystem()
	loadEfficient(system, []uint16{6, 70})

	if err := system.Step(); err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if system.ProgramCounter != 70 {
		t.Errorf("Program counter should've been 70, was actually %d\n", system.ProgramCounter)
		t.Fail()
	}
}

//Test the jump conditionals by jumping over noops and not jumping
func TestJumpConditionals(t *testing.T) {
	system := interpreter.NewSystem()
	loadEfficient(system, []uint16{opcodes.JT, 1, 5, opcodes.NOOP, opcodes.NOOP, opcodes.JT, 0, 5, opcodes.NOOP, opcodes.JF, 0, 15, opcodes.NOOP, opcodes.NOOP, opcodes.NOOP, opcodes.JF, 1, 15, 0})

	//JT 1 5; PC = 5 after this
	if err := system.Step(); err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if system.ProgramCounter != 5 {
		t.Error("Expected program counter to be 5 after JT 1 5, was actually", system.ProgramCounter)
		t.Fail()
	}

	//JT 0 5; PC = 8 after this, as it doesn't run
	if err := system.Step(); err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if system.ProgramCounter != 8 {
		t.Error("Expected program counter to be 8 after JT 0 5, was actually", system.ProgramCounter)
		t.Fail()
	}

	//NOOP
	if err := system.Step(); err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	//JF 0 15; PC = 15
	if err := system.Step(); err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if system.ProgramCounter != 15 {
		t.Error("Expected program counter to be 15 after JF 0 15, was actually", system.ProgramCounter)
		t.Fail()
	}

	//JF 1 15; PC = 18 as it doesn't run
	if err := system.Step(); err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if system.ProgramCounter != 18 {
		t.Error("Expected program counter to be 18 after JF 1 15, was actually", system.ProgramCounter)
		t.Fail()
	}
}

//Tries setting the register to a valid value, as well as an invalid value.
func TestSet(t *testing.T) {
	system := interpreter.NewSystem()
	loadEfficient(system, []uint16{opcodes.SET, 32768, 50, opcodes.SET, 32769, 32768})

	//SET 32768 50; Sets register 0 to 50
	if err := system.Step(); err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if system.Registers[0] != 50 {
		t.Error("Expected register 0 to be 50, got", system.Registers[0])
		t.Fail()
	}

	//SET 32769 32768; sets register 1 to be the value of register 0 (50)
	if err := system.Step(); err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if system.Registers[1] != 50 {
		t.Error("Expected register 1 to be 50, got", system.Registers[1])
		t.Fail()
	}
}

//Tests the conditional set operations, by giving them both true and false conditions.
func TestCondSet(t *testing.T) {
	system := interpreter.NewSystem()
	loadEfficient(system, []uint16{opcodes.EQ, 32768, 1, 1, opcodes.EQ, 32769, 1, 5, opcodes.GT, 32770, 5, 1, opcodes.GT, 32771, 1, 5})

	//eq 32768 1 1; register 0 = 1
	if err := system.Step(); err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if system.Registers[0] != 1 {
		t.Error("Expected register 0 to be 1, got ", system.Registers[0])
	}

	//eq 32769 1 5; register 1 = 0
	if err := system.Step(); err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if system.Registers[1] != 0 {
		t.Error("Expected register 1 to be 0, got ", system.Registers[1])
	}

	//gt 32770 5 1; register 2 = 1
	if err := system.Step(); err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if system.Registers[2] != 1 {
		t.Error("Expected register 2 to be 1, got ", system.Registers[2])
	}

	//gt 32771 1 5q register 3 = 0
	if err := system.Step(); err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if system.Registers[3] != 0 {
		t.Error("Expected register 3 to be 0, got ", system.Registers[3])
	}
}

//Tests the stack by pushing and popping from it
func TestStack(t *testing.T){
	system := interpreter.NewSystem()
	loadEfficient(system, []uint16{opcodes.PUSH, 5, opcodes.PUSH, 6, opcodes.POP, 32768, opcodes.POP, 32769})

	//push 5; stack len = 1, stack[0] = 5
	if err := system.Step(); err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if system.Stack.Value != 5 {
		t.Error("Expected a top stack value of 5, got ", system.Stack.Value)
		t.Fail()
	}

	//push 6; stack len = 2, stack [1] = 6
	if err := system.Step(); err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if system.Stack.Value != 6 {
		t.Error("Expected a top stack value of 6, got ", system.Stack.Value)
		t.Fail()
	}

	//pop 32768; stack len = 1, register[0] = 6
	if err := system.Step(); err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if system.Registers[0] != 6 {
		t.Error("Expected register 0 to hold 6, got", system.Registers[0])
		t.Fail()
	}

	if system.Stack.Value != 5 {
		t.Error("Expected a top stack value of 5, got ", system.Stack.Value)
		t.Fail()
	}

	//pop 32769; stack len = 0, register[1] = 5
	if err := system.Step(); err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if system.Registers[1] != 5 {
		t.Error("Expected register 1 to hold 5, got", system.Registers[1])
		t.Fail()
	}

	if system.Stack != nil {
		t.Error("Expected a top stack value of nil, got ", system.Stack)
		t.Fail()
	}
}

//Tests binary and
func TestAnd(t *testing.T) {
	system := interpreter.NewSystem()
	loadEfficient(system, []uint16{opcodes.AND, 32768, 0x00FF, 0x000F})

	//and 32768 255 15; register[0] = 15
	if err := system.Step(); err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if system.Registers[0] != 0x000F {
		t.Errorf("Expected register 0 to hold %d, got %d\n", 0x000F, system.Registers[0])
	}
}

//Tests binary or
func TestOr(t *testing.T) {
	system := interpreter.NewSystem()
	loadEfficient(system, []uint16{opcodes.OR, 32768, 0x00FF, 0x0000})

	//or 32768 255 0; register[0] = 255
	if err := system.Step(); err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if system.Registers[0] != 0x00FF {
		t.Errorf("Expected register 0 to hold %d, got %d\n", 0x00FF, system.Registers[0])
	}
}

//Tests binary not
func TestNot(t *testing.T) {
	system := interpreter.NewSystem()
	loadEfficient(system, []uint16{opcodes.NOT, 32768, 0x00FF})

	//not 32768 0x00FF; register[0] == 0x7F00
	if err := system.Step(); err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if system.Registers[0] != 0x7F00 {
		t.Errorf("Expected register 0 to hold %d, got %d\n", 0x7F00, system.Registers[0])
	}
}

//Tests the call and return system
func TestCallRet(t *testing.T) {
	system := interpreter.NewSystem()
	loadEfficient(system, []uint16{opcodes.CALL, 3, opcodes.HALT, opcodes.RET})

	//call 3; pc = 3, stack len = 1, stack 0 = 2
	if err := system.Step(); err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if system.ProgramCounter != 3 {
		t.Error("Expected program counter to jump to 3 after call 3, got", system.ProgramCounter)
		t.Fail()
	}

	if system.Stack != nil {
		if system.Stack.Value != 2 {
			t.Error("Top-stack value should be 2, got", system.Stack.Value)
			t.Fail()
		}
	} else {
		t.Error("Expected a non-empty stack.")
		t.Fail()
	}

	//ret; pc = 2, stack len = 0
	if err := system.Step(); err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if system.ProgramCounter != 2 {
		t.Error("Expected program counter to return to 2 after ret, got", system.ProgramCounter)
		t.Fail()
	}

	if system.Stack != nil {
		t.Error("Expected empty stack.")
		t.Fail()
	}
}

//Writes to memory, then reads from it
func TestMemory(t *testing.T) {
	system := interpreter.NewSystem()
	loadEfficient(system, []uint16{opcodes.WMEM, 50, 69, opcodes.RMEM, 32768, 50})

	if err := system.Step(); err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if system.Memory[50] != 69 {
		t.Error("Expected memory[50] to be 69, but got", system.Memory[50])
		t.Fail()
	}

	if err := system.Step(); err != nil {
		t.Error(err.Error())
		t.Fail()
	}

	if system.Memory[50] != 69 || system.Registers[0] != 69 {
		t.Error("Expected memory[50] and registers[0] to be 69, but got", system.Memory[50], system.Registers[0], "respectively")
	}
}