package interpreter

import "fmt"

func out(s *System) error {
	if a, err := s.GetValue(s.Memory[s.ProgramCounter+1]); err == nil {
		fmt.Printf("%c", rune(a))
	} else {
		return err
	}

	return nil
}