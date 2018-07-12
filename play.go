package play

type Play struct {
	play []interface{}
}

func New() *Play {
	return &Play{}
}

func (s *Play) Push(v interface {}) {
	s.play = append(s.play, v)
}

func (s *Play) Pop() interface{} {
	if len(s.play) == 0 {
		return nil
	} else {
		v := s.play[len(s.play)-1]
		s.play = s.play[0 : len(s.play) -1]
		return v
	}
}

func (s *Play) Top() interface{} {
	if len(s.play) == 0 {
		return nil
	} else {
		return s.play[len(s.play]) -1]
	}
}
