package interpreter

import "errors"

func pop(s *System) error {
	if a, err := s.GetLocation(s.Memory[s.ProgramCounter+1], false); err == nil || s.Stack != nil {
		*a = s.Stack.Value
		s.Stack = s.Stack.Last
	} else if err != nil{
		return err
	} else {
		return errors.New("can't pop from an empty stack")
	}

	return nil
}