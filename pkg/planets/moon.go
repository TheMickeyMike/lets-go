package planets

import "time"

// Moon is planet
type Moon struct {
	name string
}

// NewMoon creates new Moon
func NewMoon(name string) *Moon {
	return &Moon{
		name: name,
	}
}

// Landing measure landing duration
func (m *Moon) Landing() time.Duration {
	return time.Duration(time.Minute * 2)
}

func (m *Moon) String() string {
	return m.name
}
