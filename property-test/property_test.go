package property

import "testing"

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 gets I", 1, "I"},
		{"2 gets II", 2, "II"},
		{"3 gets III", 3, "III"},
		{"4 gets converted to IV", 4, "IV"},
	}

	for _, test := range cases {

		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			want := test.Want

			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})

	}

}
