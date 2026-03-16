package mastermind

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCartesianProductOrderTwo(t *testing.T) {
	assert.Equal(
		t,
		[]string{"aa", "ba", "ca", "ab", "bb", "cb", "ac", "bc", "cc"},
		cartesianProduct([]string{"abc", "abc"}))
}
