package regex

import (
	"testing"
)

func TestIsUpper(t *testing.T) {
	expected := true
	got := IsUpper("Gohex")

	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}

func TestStringToHyphen(t *testing.T) {
	expected := "go-hex"
	got := StringToHyphen("go-hex")

	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}
