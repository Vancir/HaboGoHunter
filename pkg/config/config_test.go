package config

import (
	"fmt"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		input  string
		output *Config
		err    error
	}{
		{
			"testdata/example.json",
			&Config{
				LogDir: "foobar",
			},
			nil,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			cfg, err := LoadConfig(test.input)
			if err != test.err {
				t.Fatalf("bad error: want '%v', got '%v'", test.err, err)
			}
			if test.output.LogDir != cfg.LogDir {
				t.Fatalf("bad output: want '%v', got '%v'", test.output, cfg)
			}

		})
	}
}
