package interpreter

func in(s *System) error {
	if a, err := s.GetLocation(s.Memory[s.ProgramCounter+1], false); err == nil {
		if b, err := reader.ReadByte(); err == nil {
			*a = uint16(b)
		}
	} else {
		return err
	}

	return nil
}