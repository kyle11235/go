package pk

import "testing"

func TestBar(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Bar(c.in)
		if got != c.want {
			t.Errorf("Bar(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
