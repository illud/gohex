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

func TestToKebabCase(t *testing.T) {
	expected := "go-hex"
	got := ToKebabCase("go_Hex")

	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}

func TestFormatSnakeCaseToCamelCase(t *testing.T) {
	expected := "goHex"
	got := FormatSnakeCaseToCamelCase("go_hex")

	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}

func TestDashToCamelCase(t *testing.T) {
	expected := "GoHex"
	got := DashToCamel("go-hex")

	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}

func TestCapitalizeFirstLetter(t *testing.T) {
	expected := "Gohex"
	got := CapitalizeFirstLetter("gohex")

	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}

func TestFormatHyphenToCamelCase(t *testing.T) {
	expected := "goHex"
	got := FormatHyphenToCamelCase("go-hex")

	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	}
}
