package planets

import "time"

// Planet is interface for all planets
type Planet interface {
	Landing() time.Duration
	String() string
}
