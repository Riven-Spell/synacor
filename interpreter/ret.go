package interpreter

import "errors"

func ret(s *System) error {
	if s.Stack == nil {
		s.Halted = true
		return errors.New("can't return with an empty stack")
	} else {
		s.ProgramCounter = s.Stack.Value
		s.Stack = s.Stack.Last
		s.DoJump = false
	}

	return nil
}