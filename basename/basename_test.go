package basename

import "testing"

func TestBasename(t *testing.T) {
	e := []struct {
		in       string
		expected string
	}{
		{"/a.", "a"},
		{"", ""},
		{".", ""},
		{"a.", "a"},
		{"a.b", "a"},
		{"a.b.c", "a.b"},
		{"./", ""},
		{"./a", "a"},
		{"./a.b", "a"},
		{".a", ""},
		{"/", ""},
		{"/.", ""},
		{"/.a", ""},
		{"//", ""},
		{"/a", "a"},
		{"/ab", "ab"},
		{"/ab/", ""},
		{"/a.b", "a"},
		{"/a/b.c", "b"},
		{"/a/b.c/", ""},
	}
	for _, e := range e {
		if actual := basename(e.in); e.expected != actual {
			t.Errorf("basename(%s) expected: %s, actual:%s", e.in, e.expected, actual)
		}
	}
}
