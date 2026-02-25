package validparentheses

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "empty string",
			s:    "",
			want: true,
		},
		{
			name: "basic valid",
			s:    "()[]{}",
			want: true,
		},
		{
			name: "mismatch",
			s:    "(]",
			want: false,
		},
		{
			name: "crossed",
			s:    "([)]",
			want: false,
		},
		{
			name: "nested valid",
			s:    "{[]}",
			want: true,
		},
		{
			name: "extra opening",
			s:    "(((",
			want: false,
		},
		{
			name: "extra closing",
			s:    "){",
			want: false,
		},
		{
			name: "invalid char",
			s:    "a(b)c",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValid(tt.s)
			if got != tt.want {
				t.Fatalf("IsValid(%q) = %t, want %t", tt.s, got, tt.want)
			}
		})
	}
}
