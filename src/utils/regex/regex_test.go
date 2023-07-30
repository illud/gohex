package regex

import (
	"testing"
)

func TestIsUpper(t *testing.T) {
	expected := true
	got := IsUpper("gohex")

	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}

func TestStringToHyphen(t *testing.T) {
	expected := "go-jira"
	got := StringToHyphen("gohex")

	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}
