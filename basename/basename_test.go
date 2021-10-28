package basename

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBasename(t *testing.T) {
	var tests = []struct {
		expected string
		input    string
	}{
		{"a", "/a."},
		{"", ""},
		{"", "."},
		{"a", "a."},
		{"a", "a.b"},
		{"a.b", "a.b.c"},
		{"", "./"},
		{"a", "./a"},
		{"a", "./a.bb"},
		{"", ".a"},
		{"", "/"},
		{"", "/."},
		{"", "/.a"},
		{"", "//"},
		{"a", "/a"},
		{"ab", "/ab"},
		{"", "/ab/"},
		{"a", "/a.b"},
		{"b", "/a/b.c"},
		{"", "/a/b.c/"},
		{"", "/a/b.c/"},
		{"a", "/a/b.c/a.Ω"},
		{"Ωa", "/aΩa/bΩa.Ωac/Ωa.aΩbΩc"},
	}
	for _, testCase := range tests {
		actual := basename(testCase.input)
		require.Equal(t, testCase.expected, actual, "basename(%s) expected: %s, actual:%s", testCase.input, testCase.expected, actual)
	}
}
