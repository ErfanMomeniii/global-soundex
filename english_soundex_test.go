package globalsoundex

import "testing"

func TestEnglishSoundex_Encode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Robert", "r163"},
		{"Rupert", "r163"},
		{"Ashcraft", "a261"},
		{"Honeyman", "h500"},
		{"H123!ney?m?!an", "h500"},
	}

	s := newEnglishSoundex()

	for _, test := range tests {
		output := s.Encode(test.input)

		if output != test.expected {
			t.Fatalf("expected %v, got %v", test.expected, output)
		}
	}
}

func TestEnglishSoundex_AddEntities_And_Suggest(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Robert", "Rob"},
		{"Rupert", "Rob"},
		{"Ashcraft", "AirCraft"},
		{"Honeyman", "Honey"},
		{"H123!ney?m?!an", "Honey"},
		{"travell", "Travel"},
	}

	s := newEnglishSoundex()
	s.AddEntities([]string{"Rob", "AirCraft", "Honey", "Travel"})

	for _, test := range tests {
		output := s.Suggest(test.input)

		if output != test.expected {
			t.Fatalf("expected %v, got %v", test.expected, output)
		}
	}
}
