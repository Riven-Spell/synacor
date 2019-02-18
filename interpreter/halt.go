package interpreter

func halt(s *System) error {
	s.Halted = true
	return nil
}