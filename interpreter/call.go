package interpreter

import "synacor/interpreter/opcodes"

func call(s *System) error {
	if a, err := s.GetValue(s.Memory[s.ProgramCounter+1]); err == nil {
		s.Stack = append(s.Stack, s.ProgramCounter + opcodes.OpcodeLength[s.Memory[s.ProgramCounter]])
		s.ProgramCounter = a
		s.DoJump = false
	} else {
		return err
	}

	return nil
}