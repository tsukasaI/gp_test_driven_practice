package prop

import "testing"

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 gets converted I", 1, "I"},
		{"2 gets converted II", 2, "II"},
		{"3 gets converted III", 3, "III"},
		{"4 gets converted IV", 4, "IV"},
		{"5 gets converted V", 5, "V"},
		{"6 gets converted VI", 6, "VI"},
		{"7 gets converted VII", 7, "VII"},
		{"9 gets converted IX", 9, "IX"},
		{"40 gets converted to XL", 40, "XL"},
		{"47 gets converted to XLVII", 47, "XLVII"},
		{"49 gets converted to XLIX", 49, "XLIX"},
		{"50 gets converted to L", 50, "L"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			want := test.Want

			if got != want {
				t.Errorf("want %q, got %q", want, got)
			}
		})
	}
}
