package interpreter

func jt(s *System) error {
	if val, err := s.GetValue(s.Memory[s.ProgramCounter+1]); err == nil {
		if val != 0 {
			if s.ProgramCounter, err = s.GetValue(s.Memory[s.ProgramCounter+2]); err != nil {
				return err
			}
			s.DoJump = false
		}
	} else {
		return err
	}

	return nil
}