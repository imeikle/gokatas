package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEcho(t *testing.T) {
	tests := []struct {
		newline bool
		sep     string
		args    []string
		want    string
	}{
		{false, "", []string{}, ""},
		{true, " ", []string{}, "\n"},
		{true, "\t", []string{"one", "two", "three"}, "one\ttwo\tthree\n"},
		{false, ":", []string{"1", "2", "3"}, "1:2:3"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("echo(%v, %q, %q)", test.newline, test.sep, test.args)

		out = new(bytes.Buffer) // captured output

		if err := echo(test.newline, test.sep, test.args); err != nil {
			t.Errorf("%s failed: %v", descr, err)
			continue
		}

		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
