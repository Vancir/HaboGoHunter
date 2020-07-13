package pathutil

import (
	"fmt"
	"testing"
)

func TestWithSuffix(t *testing.T) {
	tests := []struct {
		input  string
		suffix string
		output string
	}{
		{
			input:  "/home/foobar/example.txt",
			suffix: ".json",
			output: "/home/foobar/example.json",
		},
		{
			input:  "/home/foobar/example",
			suffix: ".txt",
			output: "/home/foobar/example.txt",
		},
		{
			input:  "/home/foobar/",
			suffix: ".txt",
			output: "/home/foobar.txt",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			result := WithSuffix(test.input, test.suffix)
			if result != test.output {
				t.Fatalf("bad output: want '%v', got '%v'", test.output, result)
			}
		})
	}
}
