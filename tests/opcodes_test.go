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

func TestJumpConditionals(t *testing.T) {
	system := interpreter.NewSystem()
	loadEfficient(system, []uint16{opcodes.JT, 1, 5, opcodes.NOOP, opcodes.NOOP, opcodes.JT, 0, 5, opcodes.NOOP, opcodes.JF, 0, 15, opcodes.NOOP, opcodes.NOOP, opcodes.NOOP, opcodes.JF, 1, 15, 0})


}