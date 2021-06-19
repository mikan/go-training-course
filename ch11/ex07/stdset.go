package intset

type StdSet struct {
	set map[int]struct{}
}

func (s *StdSet) Has(x int) bool {
	if s.set == nil {
		return false
	}
	_, ok := s.set[x]
	return ok
}

func (s *StdSet) Add(x int) {
	if s.set == nil {
		s.set = make(map[int]struct{})
	}
	s.set[x] = struct{}{}
}

func (s *StdSet) UnionWith(t *StdSet) {
	for k := range t.set {
		s.set[k] = struct{}{}
	}
}
