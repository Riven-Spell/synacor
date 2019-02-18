package interpreter

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"synacor/interpreter/opcodes"
)

const nMod = 32768

var reader = bufio.NewReader(os.Stdin)

var funcMap = map[uint16]func(*System)error{
	opcodes.HALT: halt,
	opcodes.SET: set,
	opcodes.PUSH: push,
	opcodes.POP: pop,
	opcodes.EQ: eq,
	opcodes.GT: gt,
	opcodes.JMP: jmp,
	opcodes.JT: jt,
	opcodes.JF: jf,
	opcodes.ADD: add,
	opcodes.MULT: mult,
	opcodes.MOD: mod,
	opcodes.AND: and,
	opcodes.OR: or,
	opcodes.NOT: not,
	opcodes.RMEM: rmem,
	opcodes.WMEM: wmem,
	opcodes.NOOP: noop,
	opcodes.OUT: out,
	opcodes.RET: ret,
	opcodes.CALL: call,
	opcodes.IN: in,
}

type System struct {
	Memory         [32768]uint16
	Registers      [8]uint16
	Stack          []uint16
	ProgramCounter uint16
	Halted         bool
	DoJump         bool
}

func NewSystem() *System {
	return &System{
		Memory:         [32768]uint16{},
		Registers:      [8]uint16{},
		Stack:          make([]uint16, 0),
		ProgramCounter: 0,
	}
}

func (s *System) LoadProgram(program []byte) {
	for i := 0; i < len(program); i += 2 {
		toMem := uint16(program[i+1])
		toMem = toMem << 8
		toMem = toMem | uint16(program[i])

		s.Memory[i/2] = toMem
	}
}

func (s *System) StartSystem() error {
	fmt.Println("Starting execution")

	for !s.Halted {
		if err := s.Step(); err != nil {
			s.Halted = true
			return err
		}
	}

	return nil
}

func (s *System) GetLocation(location uint16, memPtrPossible bool) (*uint16, error) {
	if location <= 32767 && memPtrPossible {
		return &s.Memory[location], nil
	} else if location <= 32775 {
		return &s.Registers[location-32768], nil
	}

	return nil, errors.New(strconv.Itoa(int(location)) + " is a invalid location")
}

func (s *System) GetValue(value uint16) (uint16, error) {
	if value <= 32767 {
		return value, nil
	} else if value <= 32775 {
		return s.Registers[value-32768], nil
	}

	return 0, errors.New(strconv.Itoa(int(value)) + " is a invalid value")
}

func (s *System) Step() error {
	if s.Halted {
		return errors.New("CPU is halted, cannot continue")
	}

	s.DoJump = true

	op := s.Memory[s.ProgramCounter]

	toIncrement := opcodes.OpcodeLength[op]
	if err := funcMap[op](s); err != nil {
		return err
	}

	if s.DoJump {
		s.ProgramCounter += toIncrement
	}

	return nil
}
