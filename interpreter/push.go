package interpreter

func push(s *System) error {
	if a, err := s.GetValue(s.Memory[s.ProgramCounter+1]); err == nil {
		s.Stack = append(s.Stack, a)
	} else {
		return err
	}

	return nil
}
