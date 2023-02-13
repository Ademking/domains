package main

import (
	"testing"
)

func TestExtractDomain(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"http://www.example.com", "www.example.com"},
		{"https://subdomain.example.co.uk/path", "subdomain.example.co.uk"},
		{"invalid url", ""},
	}

	for _, tc := range testCases {
		actual := extract_domain(tc.input)
		if actual != tc.expected {
			t.Errorf("expected %s, but got %s for input %s", tc.expected, actual, tc.input)
		}
	}
}

func TestFormatDomain(t *testing.T) {
	testCases := []struct {
		prefix   string
		domain   string
		suffix   string
		expected string
	}{
		{"https://", "www.example.com", "", "https://www.example.com"},
		{"", "subdomain.example.co.uk", "/path", "subdomain.example.co.uk/path"},
	}

	for _, tc := range testCases {
		actual := format_domain(tc.prefix, tc.domain, tc.suffix)
		if actual != tc.expected {
			t.Errorf("expected %s, but got %s for prefix %s, domain %s, suffix %s", tc.expected, actual, tc.prefix, tc.domain, tc.suffix)
		}
	}
}
