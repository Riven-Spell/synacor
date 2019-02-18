package interpreter

import "errors"

func pop(s *System) error {
	if a, err := s.GetLocation(s.Memory[s.ProgramCounter+1], false); err == nil || len(s.Stack) == 0 {
		*a = s.Stack[len(s.Stack) - 1]
		s.Stack = s.Stack[:len(s.Stack) - 1]
	} else if err != nil{
		return err
	} else {
		return errors.New("can't pop from an empty stack")
	}

	return nil
}