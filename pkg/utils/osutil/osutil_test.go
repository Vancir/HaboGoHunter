package osutil

import (
	"fmt"

	"testing"
)

func TestRunCmd(t *testing.T) {
	t.Helper()

	tests := []struct {
		timeout int64
		dir     string
		bin     string
		args    string
		output  string
		err     error
	}{
		{
			1, ".", "echo", "foobar", "foobar", nil,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			result, err := RunCmd(test.timeout, test.dir, test.bin, test.args)
			if err != test.err {
				t.Fatalf("bad error: want '%v', got '%v'", test.err, err)
			}
			if result != test.output {
				t.Fatalf("bad output: want '%v', got '%v'", test.output, result)
			}
		})
	}
}
