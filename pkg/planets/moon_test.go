package planets_test

import (
	"testing"
	"time"

	"github.com/TheMickeyMike/lets-go/pkg/planets"
	"github.com/stretchr/testify/assert"
)

func TestTimeConsuming(t *testing.T) {
	moon := planets.NewMoon("TestMoon")
	landingTime := moon.Landing()
	assert.Equal(t, time.Minute*2, landingTime)

}
