package interpreter

func rmem(s *System) error {
	if a, err := s.GetLocation(s.Memory[s.ProgramCounter+1], true); err == nil {
		if b, err := s.GetValue(s.Memory[s.ProgramCounter+2]); err == nil {
			*a = s.Memory[b]
		} else {
			return err
		}
	} else {
		return err
	}

	return nil
}
