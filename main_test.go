package main

import (
	"reflect"
	"testing"
)

func TestExtractDomain(t *testing.T) {
	input := "example.com\nexample.org\nnot a domain"
	want := []string{"example.com", "example.org"}

	got := extract_domain(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("extract_domain(%q) = %q; want %q", input, got, want)
	}
}
