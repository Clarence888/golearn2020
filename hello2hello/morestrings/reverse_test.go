package morestrings

import "testing"

func TestReverseRunes(t *testing.T) {
	cases := []struct{
		in, want string
	}{
		{"HelloWorld","dlroWolleH"},
		{"nihao","oahin"},
		{"",""},
		{"123","321"},
	}
	for _,c := range cases{
		got := ReverseRunes(c.in)
		if got != c.want {
			t.Errorf("ReverseRunes(%q) == %q,want %q", c.in,got, c.want)
		}
	}
}
