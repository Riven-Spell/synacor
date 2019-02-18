package interpreter

func eq(s *System) error {
	if a, err := s.GetLocation(s.Memory[s.ProgramCounter+1], false); err == nil {
		if b, err := s.GetValue(s.Memory[s.ProgramCounter+2]); err == nil {
			if c, err := s.GetValue(s.Memory[s.ProgramCounter+3]); err == nil {
				if b == c {
					*a = 1
				} else {
					*a = 0
				}
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