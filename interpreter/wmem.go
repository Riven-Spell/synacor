package interpreter

func wmem(s *System) error {
	if al, err := s.GetValue(s.Memory[s.ProgramCounter+1]); err == nil {
		if a, err := s.GetLocation(al, true); err == nil {
			if b, err := s.GetValue(s.Memory[s.ProgramCounter+2]); err == nil {
				*a = b
			} else {
				return err
			}
		} else {
			return err
		}
	} else {
		return err
	}

	return nil
}