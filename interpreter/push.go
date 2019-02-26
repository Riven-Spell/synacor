package interpreter

func push(s *System) error {
	if a, err := s.GetValue(s.Memory[s.ProgramCounter+1]); err == nil {
		s.Stack = &LinkedUInt16{
			Last: s.Stack,
			Value: a,
		}
	} else {
		return err
	}

	return nil
}
