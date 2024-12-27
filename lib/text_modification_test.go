package lib

import (
	"testing"

	"github.com/gofiber/fiber/v2/utils"
)

func TestTrimSpace(t *testing.T) {
	tests := []string{
		"EXAMPLE",
		"EXAMPLE ",
		"EXAMPLE  ",
		" EXAMPLE",
		"  EXAMPLE",
		" EXAMPLE ",
		"  EXAMPLE  ",
	}

	for _, test := range tests {
		result := TrimSpace(test)
		utils.AssertEqual(t, "EXAMPLE", result)

		resultPtr := TrimSpacePtr(&test)
		utils.AssertEqual(t, "EXAMPLE", *resultPtr)
	}
}

func TestTitleCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"2 PIECES UP TO 23 KG EACH", "2 Pieces Up To 23 Kg Each"},
		// Add more test cases as needed
	}

	for _, test := range tests {
		actual := TitleCase(test.input)
		expected := test.expected

		if actual != expected {
			t.Errorf("TitleCase(%s): expected '%s', got '%s'", test.input, expected, actual)
		}
	}
}

func TestSentenceCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"2 PIECES UP TO 23 KG EACH", "2 pieces up to 23 kg each"},
		// Add more test cases as needed
	}

	for _, test := range tests {
		actual := SentenceCase(test.input)
		expected := test.expected

		if actual != expected {
			t.Errorf("SentenceCase(%s): expected '%s', got '%s'", test.input, expected, actual)
		}
	}
}
