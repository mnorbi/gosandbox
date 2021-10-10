package basename

import "testing"

func TestBasename(t *testing.T) {
	// TODO Just a pro tip: start using testify/require or testify/assert
	// import "github.com/stretchr/testify/require"
	// require.Equal(t, "a", basename("/a.")
	// require.Equal(t, "", basename("")
	// ...
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
