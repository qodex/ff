package ff

import (
	"fmt"
	"testing"
)

func TestHash9(t *testing.T) {
	for i := range 1000 {
		hash := Hash9(fmt.Sprintf("%d%b", i, byte(i)))
		if len(fmt.Sprintf("%d", hash)) != 9 {
			t.Errorf("not len 9: %d", hash)
		}
	}
}
