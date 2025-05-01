package globalsoundex

import "testing"

func TestArabicSoundex_Encode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"عبدالرحمن", "A465"},
		{"ابويوسف", "Y210"},
		{"محمد", "M530"},
		{"محمد123", "M530"},
	}

	s := newArabicSoundex()

	for _, test := range tests {
		output := s.Encode(test.input)

		if output != test.expected {
			t.Fatalf("expected %v, got %v", test.expected, output)
		}
	}
}

func TestArabicSoundex_AddEntities_And_Suggest(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"الحسین", "حسین"},
		{"الحسی", "حسین"},
		{"المنصور", "منصور"},
		{"مهممد", "محمد"},
	}

	s := newArabicSoundex()
	s.AddEntities([]string{"محمد", "منصور", "احمد", "حسین"})

	for _, test := range tests {
		output := s.Suggest(test.input)

		if output != test.expected {
			t.Fatalf("expected %v, got %v", test.expected, output)
		}
	}
}
