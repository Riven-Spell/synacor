package interpreter

import "errors"

func ret(s *System) error {
	if len(s.Stack) == 0 {
		s.Halted = true
		return errors.New("can't return with an empty stack")
	} else {
		s.ProgramCounter = s.Stack[len(s.Stack) - 1]
		s.Stack = s.Stack[:len(s.Stack) - 1]
		s.DoJump = false
	}

	return nil
}