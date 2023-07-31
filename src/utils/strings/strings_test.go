package strings

import (
	"testing"
)

func TestGetFirstCharacterOfString(t *testing.T) {
	expected := "g"
	got := GetFirstCharacterOfString("gohex")

	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}
