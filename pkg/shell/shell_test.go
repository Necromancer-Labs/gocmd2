package shell

import (
	"testing"

	"github.com/google/shlex"
)

func TestShlexParsing(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "simple command",
			input:    "echo hello",
			expected: []string{"echo", "hello"},
		},
		{
			name:     "quoted argument",
			input:    `echo "hello world"`,
			expected: []string{"echo", "hello world"},
		},
		{
			name:     "single quoted argument",
			input:    `echo 'hello world'`,
			expected: []string{"echo", "hello world"},
		},
		{
			name:     "mixed quotes",
			input:    `cmd "arg one" 'arg two' three`,
			expected: []string{"cmd", "arg one", "arg two", "three"},
		},
		{
			name:     "escaped quote in double quotes",
			input:    `echo "say \"hello\""`,
			expected: []string{"echo", `say "hello"`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args, err := shlex.Split(tt.input)
			if err != nil {
				t.Fatalf("shlex.Split(%q) returned error: %v", tt.input, err)
			}

			if len(args) != len(tt.expected) {
				t.Fatalf("shlex.Split(%q) = %v, want %v", tt.input, args, tt.expected)
			}

			for i, arg := range args {
				if arg != tt.expected[i] {
					t.Errorf("shlex.Split(%q)[%d] = %q, want %q", tt.input, i, arg, tt.expected[i])
				}
			}
		})
	}
}
