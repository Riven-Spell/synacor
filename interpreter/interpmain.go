package interpreter

import (
	"errors"
	"fmt"
	"strconv"
	"synacor/interpreter/opcodes"
)

const nMod = 32768

type System struct {
	Memory         [32768]uint16
	Registers      [8]uint16
	Stack          []uint16
	ProgramCounter uint16
	Halted         bool
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

	switch s.Memory[s.ProgramCounter] {
	case opcodes.NOOP: //Nothing happens.

	case opcodes.HALT: //The system halts.
		s.Halted = true

	case opcodes.ADD: //Add b & c, store to a; add a b c
		if a, err := s.GetLocation(s.Memory[s.ProgramCounter+1], false); err == nil {
			if b, err := s.GetValue(s.Memory[s.ProgramCounter+2]); err == nil {
				if c, err := s.GetValue(s.Memory[s.ProgramCounter+3]); err == nil {
					*a = b + c
					*a %= 32768
				} else {
					return err
				}
			} else {
				return err
			}
		} else {
			return err
		}

	case opcodes.OUT: //Output the ASCII character value of a; out a
		if a, err := s.GetValue(s.Memory[s.ProgramCounter+1]); err == nil {
			fmt.Printf("%c", rune(a))
		} else {
			return err
		}

	case opcodes.JMP: //Unconditionally jump to a; jmp a
		var err error
		if s.ProgramCounter, err = s.GetValue(s.Memory[s.ProgramCounter+1]); err != nil {
			return err
		} else {
			return nil //Bypass program counter increment.
		}

	case opcodes.JT: //Jump to b if a is nonzero; jt a b
		if val, err := s.GetValue(s.Memory[s.ProgramCounter+1]); err == nil {
			if val != 0 {
				if s.ProgramCounter, err = s.GetValue(s.Memory[s.ProgramCounter+2]); err != nil {
					return err
				} else {
					return nil //Bypass program encounter increment.
				}
			}
		} else {
			return err
		}

	case opcodes.JF: //Jump to b if a is zero; jf a b
		if val, err := s.GetValue(s.Memory[s.ProgramCounter+1]); err == nil {
			if val == 0 {
				if s.ProgramCounter, err = s.GetValue(s.Memory[s.ProgramCounter+2]); err != nil {
					return err
				} else {
					return nil
				}
			}
		} else {
			return err
		}

	case opcodes.SET:
		if a, err := s.GetLocation(s.Memory[s.ProgramCounter+1], false); err == nil {
			if b, err := s.GetValue(s.Memory[s.ProgramCounter+2]); err == nil {
				*a = b
			} else {
				return err
			}
		} else {
			return err
		}

	case opcodes.EQ:
		if a, err := s.GetLocation(s.Memory[s.ProgramCounter+1], false); err == nil {
			if b, err := s.GetValue(s.Memory[s.ProgramCounter+2]); err == nil {
				if c, err := s.GetValue(s.Memory[s.ProgramCounter+3]); err == nil {
					if b == c {
						*a = 1
					} else {
						*a = 0
					}
				} else {
					return err
				}
			} else {
				return err
			}
		} else {
			return err
		}

	case opcodes.GT:
		if a, err := s.GetLocation(s.Memory[s.ProgramCounter+1], false); err == nil {
			if b, err := s.GetValue(s.Memory[s.ProgramCounter+2]); err == nil {
				if c, err := s.GetValue(s.Memory[s.ProgramCounter+3]); err == nil {
					if b > c {
						*a = 1
					} else {
						*a = 0
					}
				} else {
					return err
				}
			} else {
				return err
			}
		} else {
			return err
		}

	case opcodes.PUSH:
		if a, err := s.GetValue(s.Memory[s.ProgramCounter+1]); err == nil {
			s.Stack = append(s.Stack, a)
		} else {
			return err
		}

	case opcodes.POP:
		if a, err := s.GetLocation(s.Memory[s.ProgramCounter+1], false); err == nil || len(s.Stack) == 0 {
			*a = s.Stack[len(s.Stack) - 1]
			s.Stack = s.Stack[:len(s.Stack) - 1]
		} else if err != nil{
			return err
		} else {
			return errors.New("can't pop from an empty stack")
		}

	case opcodes.AND:
		if a, err := s.GetLocation(s.Memory[s.ProgramCounter+1], false); err == nil {
			if b, err := s.GetValue(s.Memory[s.ProgramCounter+2]); err == nil {
				if c, err := s.GetValue(s.Memory[s.ProgramCounter+3]); err == nil {
					*a = b & c
				} else {
					return err
				}
			} else {
				return err
			}
		} else {
			return err
		}

	case opcodes.OR:
		if a, err := s.GetLocation(s.Memory[s.ProgramCounter+1], false); err == nil {
			if b, err := s.GetValue(s.Memory[s.ProgramCounter+2]); err == nil {
				if c, err := s.GetValue(s.Memory[s.ProgramCounter+3]); err == nil {
					*a = b | c
				} else {
					return err
				}
			} else {
				return err
			}
		} else {
			return err
		}

	case opcodes.NOT:
		if a, err := s.GetLocation(s.Memory[s.ProgramCounter+1], false); err == nil {
			if b, err := s.GetValue(s.Memory[s.ProgramCounter+2]); err == nil {
				*a = ^b & 0x7FFF
			} else {
				return err
			}
		} else {
			return err
		}
	}

	s.ProgramCounter += opcodes.OpcodeLength[s.Memory[s.ProgramCounter]]

	return nil
}
