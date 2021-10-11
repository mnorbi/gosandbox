package format_numbers

import "testing"

func TestFormatNumberStr(t *testing.T) {
	e := []struct {
		in       string
		expected string
	}{
		{"-1", "-1"},
		{"-12", "-12"},
		{"-123", "-123"},
		{"-1234", "-1,234"},
		{"", ""},
		{"1", "1"},
		{"123", "123"},
		{"1234", "1,234"},
		{"12345", "12,345"},
		{"1234567", "1,234,567"},
		{"12345678", "12,345,678"},
	}
	for _, e := range e {
		if actual := formatNumberStr(e.in); e.expected != actual {
			t.Errorf("format(%s) expected: %s, actual:%s", e.in, e.expected, actual)
		}
	}
}

func TestFormatNumber(t *testing.T) {
	e := []struct {
		in       int
		expected string
	}{
		{0, "0"},
		{1, "1"},
		{123, "123"},
		{1234, "1,234"},
		{12345, "12,345"},
		{1234567, "1,234,567"},
		{12345678, "12,345,678"},
		{-1, "-1"},
		{-12, "-12"},
		{-123, "-123"},
		{-1234, "-1,234"},
	}
	for _, e := range e {
		if actual := formatNumber(e.in); e.expected != actual {
			t.Errorf("format(%d) expected: %s, actual:%s", e.in, e.expected, actual)
		}
	}
}
