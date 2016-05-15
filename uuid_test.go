package uuid

import (
	"regexp"
	"testing"
)

func TestNew(t *testing.T) {
	rgx := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	for i := 0; i < 100000; i++ {
		uuid := NewV4()
		if !rgx.MatchString(uuid) {
			t.Fatal("Invalid UUID")
		}
	}
}
