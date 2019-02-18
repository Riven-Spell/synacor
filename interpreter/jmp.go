package interpreter

func jmp(s *System) error {
	var err error
	if s.ProgramCounter, err = s.GetValue(s.Memory[s.ProgramCounter+1]); err != nil {
		return err
	}

	return nil
}